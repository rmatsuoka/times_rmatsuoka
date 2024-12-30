package api

import (
	"net/http"

	"github.com/rmatsuoka/times_rmatsuoka/internal/usercmd"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
)

func (api *API) createUser(w http.ResponseWriter, req *http.Request, user *usercmd.Creating) {
	id, err := usercmd.Create(req.Context(), api.DB, user)
	writeResult(w, struct {
		ID users.ID `json:"id"`
	}{ID: id}, err)
}

func (api *API) getUser(w http.ResponseWriter, req *http.Request) {
	userCode := req.PathValue("userCode")
	u, err := usercmd.Get(req.Context(), api.DB, userCode)
	writeResult(w, struct {
		User *users.User `json:"user"`
	}{User: u}, err)
}

func (api *API) deleteUser(w http.ResponseWriter, req *http.Request) {
	userCode := req.PathValue("userCode")
	err := usercmd.Delete(req.Context(), api.DB, userCode)
	writeResult(w, empty, err)
}

func (api *API) updateUser(w http.ResponseWriter, req *http.Request, user *usercmd.Creating) {
	userCode := req.PathValue("userCode")
	err := usercmd.Update(req.Context(), api.DB, userCode, user)
	writeResult(w, empty, err)
}
