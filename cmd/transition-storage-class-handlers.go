package cmd

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minio/minio/cmd/logger"
	iampolicy "github.com/minio/minio/pkg/iam/policy"
	"github.com/minio/minio/pkg/madmin"
)

var (
	// error returned when remote tier already exists
	errTierAlreadyExists = AdminError{
		Code:       "XMinioAdminTierAlreadyExists",
		Message:    "Specified remote tier already exists",
		StatusCode: http.StatusConflict,
	}
	// error returned when remote tier is not found
	errTierNotFound = AdminError{
		Code:       "XMinioAdminTierNotFound",
		Message:    "Specified remote tier was not found",
		StatusCode: http.StatusNotFound,
	}
)

func (api adminAPIHandlers) AddStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "AddStorageClass")

	defer logger.AuditLog(w, r, "AddStorageClass", mustGetClaimsFromToken(r))

	objectAPI, cred := validateAdminUsersReq(ctx, w, r, iampolicy.SetStorageClassAction)
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfigMgr == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	password := cred.SecretKey
	reqBytes, err := madmin.DecryptData(password, io.LimitReader(r.Body, r.ContentLength))
	if err != nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErrWithErr(ErrAdminConfigBadJSON, err), r.URL)
		return
	}

	var cfg madmin.TierConfig
	if err := json.Unmarshal(reqBytes, &cfg); err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}

	err = globalTransitionStorageClassConfigMgr.Add(cfg)
	if err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}

	err = saveGlobalTransitionStorageClassConfig()
	if err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}
	globalNotificationSys.LoadTransitionStorageClassConfig(ctx)

	writeSuccessNoContent(w)
}

func (api adminAPIHandlers) RemoveStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "RemoveStorageClass")

	defer logger.AuditLog(w, r, "RemoveStorageClass", mustGetClaimsFromToken(r))

	objectAPI, _ := validateAdminUsersReq(ctx, w, r, iampolicy.RemoveStorageClassAction)
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfigMgr == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	var vars = mux.Vars(r)
	scName := vars["storageclass"]

	globalTransitionStorageClassConfigMgr.RemoveStorageClass(scName)
	err := saveGlobalTransitionStorageClassConfig()
	if err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}
	globalNotificationSys.LoadTransitionStorageClassConfig(ctx)

	writeSuccessNoContent(w)
}

func (api adminAPIHandlers) ListStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "ListStorageClass")

	defer logger.AuditLog(w, r, "ListStorageClass", mustGetClaimsFromToken(r))

	objectAPI, _ := validateAdminUsersReq(ctx, w, r, iampolicy.ListStorageClassAction)
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfigMgr == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	storageClasses := globalTransitionStorageClassConfigMgr.ListStorageClasses()
	data, err := json.Marshal(storageClasses)
	if err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}

	writeSuccessResponseJSON(w, data)
}

func (api adminAPIHandlers) EditStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "EditStorageClass")

	defer logger.AuditLog(w, r, "EditStorageClass", mustGetClaimsFromToken(r))

	objectAPI, cred := validateAdminUsersReq(ctx, w, r, iampolicy.SetStorageClassAction)
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfigMgr == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}
	vars := mux.Vars(r)
	scName := vars["storageclass"]

	password := cred.SecretKey
	reqBytes, err := madmin.DecryptData(password, io.LimitReader(r.Body, r.ContentLength))
	if err != nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErrWithErr(ErrAdminConfigBadJSON, err), r.URL)
		return
	}

	var creds madmin.TierCreds
	if err := json.Unmarshal(reqBytes, &creds); err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}

	if err := globalTransitionStorageClassConfigMgr.Edit(scName, creds); err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}

	if err := saveGlobalTransitionStorageClassConfig(); err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}
	globalNotificationSys.LoadTransitionStorageClassConfig(ctx)

	writeSuccessNoContent(w)
}
