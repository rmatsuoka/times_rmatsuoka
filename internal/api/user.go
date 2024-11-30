package api

import (
	"net/http"

	"github.com/rmatsuoka/times_rmatsuoka/internal/usercmd"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xhttp"
)

func (api *API) createUser(w http.ResponseWriter, req *http.Request, user *usercmd.Creating) {
	id, err := usercmd.Create(req.Context(), api.DB, user)
	if err != nil {
		xhttp.WriteJSON(w, 500, map[string]any{"message": err.Error()})
		return
	}
	xhttp.WriteJSON(w, 200, map[string]any{"id": id})
}

func (api *API) getUser(w http.ResponseWriter, req *http.Request) {
	userCode := req.PathValue("userCode")
	u, err := usercmd.Get(req.Context(), api.DB, userCode)
	if err != nil {
		xhttp.WriteJSON(w, 500, map[string]any{"message": err.Error()})
		return
	}
	xhttp.WriteJSON(w, 200, map[string]any{"user": u})
}
