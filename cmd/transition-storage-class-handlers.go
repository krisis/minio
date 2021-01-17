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
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfig == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	var vars = mux.Vars(r)

	scType := vars["type"]

	switch scType {
	case "s3":
		var sc madmin.TransitionStorageClassS3
		if err := json.NewDecoder(r.Body).Decode(&sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
		if err := globalTransitionStorageClassConfig.Add(sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
	case "azure":
		var sc madmin.TransitionStorageClassAzure
		if err := json.NewDecoder(r.Body).Decode(&sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
		if err := globalTransitionStorageClassConfig.Add(sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
	case "gcs":
		var sc madmin.TransitionStorageClassGCS
		if err := json.NewDecoder(r.Body).Decode(&sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
		if err := globalTransitionStorageClassConfig.Add(sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
	}
	globalNotificationSys.LoadTransitionStorageClassConfig(ctx)
}

func (api adminAPIHandlers) RemoveStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "RemoveStorageClass")

	defer logger.AuditLog(w, r, "RemoveStorageClass", mustGetClaimsFromToken(r))

	objectAPI := newObjectLayerFn()
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfig == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	var vars = mux.Vars(r)
	scName := vars["name"]

	globalTransitionStorageClassConfig.RemoveStorageClass(scName)
	globalNotificationSys.LoadTransitionStorageClassConfig(ctx)
}

func (api adminAPIHandlers) ListStorageClassHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "RemoveStorageClass")

	defer logger.AuditLog(w, r, "RemoveStorageClass", mustGetClaimsFromToken(r))

	objectAPI := newObjectLayerFn()
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfig == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	b, err := globalTransitionStorageClassConfig.Byte()
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
	if objectAPI == nil || globalNotificationSys == nil || globalTransitionStorageClassConfig == nil {
		writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
		return
	}

	var vars = mux.Vars(r)

	scType := vars["type"]

	switch scType {
	case "s3":
		var sc madmin.TransitionStorageClassS3
		if err := json.NewDecoder(r.Body).Decode(&sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
		if err := globalTransitionStorageClassConfig.Edit(sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
	case "azure":
		var sc madmin.TransitionStorageClassAzure
		if err := json.NewDecoder(r.Body).Decode(&sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
		if err := globalTransitionStorageClassConfig.Edit(sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
	case "gcs":
		var sc madmin.TransitionStorageClassGCS
		if err := json.NewDecoder(r.Body).Decode(&sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
		if err := globalTransitionStorageClassConfig.Edit(sc); err != nil {
			writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
			return
		}
	}
	globalNotificationSys.LoadTransitionStorageClassConfig(ctx)
}
