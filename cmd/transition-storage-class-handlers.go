package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minio/minio/cmd/logger"
	"github.com/minio/minio/pkg/madmin"
)

func (api adminAPIHandlers) AddStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "AddStorageClass")

	defer logger.AuditLog(w, r, "AddStorageClass", mustGetClaimsFromToken(r))

	objectAPI := newObjectLayerFn()
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfigMgr == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	var cfg madmin.TransitionStorageClassConfig
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}
	globalNotificationSys.LoadTransitionStorageClassConfig(ctx)
}

func (api adminAPIHandlers) RemoveStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "RemoveStorageClass")

	defer logger.AuditLog(w, r, "RemoveStorageClass", mustGetClaimsFromToken(r))

	objectAPI := newObjectLayerFn()
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfigMgr == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	var vars = mux.Vars(r)
	scName := vars["name"]

	globalTransitionStorageClassConfigMgr.RemoveStorageClass(scName)
	globalNotificationSys.LoadTransitionStorageClassConfig(ctx)
}

func (api adminAPIHandlers) ListStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "ListStorageClass")

	defer logger.AuditLog(w, r, "ListStorageClass", mustGetClaimsFromToken(r))

	objectAPI := newObjectLayerFn()
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfigMgr == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	b, err := globalTransitionStorageClassConfigMgr.Bytes()
	if err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}
	w.Write(b)
}

func (api adminAPIHandlers) EditStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "EditStorageClass")

	defer logger.AuditLog(w, r, "EditStorageClass", mustGetClaimsFromToken(r))

	objectAPI := newObjectLayerFn()
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfigMgr == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	// FIXME: edit should allow only creds to be updated
	var sc madmin.TransitionStorageClassConfig
	if err := json.NewDecoder(r.Body).Decode(&sc); err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}
	if err := globalTransitionStorageClassConfigMgr.Edit(sc); err != nil {
		writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
		return
	}
	globalNotificationSys.LoadTransitionStorageClassConfig(ctx)
}
