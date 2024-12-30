package api

import (
	"net/http"

	"github.com/rmatsuoka/times_rmatsuoka/internal/usercmd"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
)

func (api *API) createUser(w http.ResponseWriter, req *http.Request, user *usercmd.Creating) {
	id, err := usercmd.Create(req.Context(), api.DB, user)
	writeResult(w, struct {
		ID string `json:"id"`
	}{ID: id.String()}, err)
}

func (api *API) getUser(w http.ResponseWriter, req *http.Request) {
	userCode := req.PathValue("userCode")
	u, err := usercmd.Get(req.Context(), api.DB, userCode)
	writeResult(w, struct {
		User *users.User `json:"user"`
	}{User: u}, err)
}
