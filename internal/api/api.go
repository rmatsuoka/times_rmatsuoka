package api

import (
	"net/http"

	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xhttp"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

type API struct {
	DB xsql.DB
}

func (api *API) Install(handle func(string, http.Handler)) {
	handle("POST /api/users", xhttp.JSONHandler(api.createUser))
	handle("GET /api/users/{userCode}", http.HandlerFunc(api.getUser))
	handle("POST /api/users/{userCode}", xhttp.JSONHandler(api.updateUser))
	handle("DELETE /api/users/{userCode}", http.HandlerFunc(api.deleteUser))

	handle("POST /api/channels", xhttp.JSONHandler(api.createChannel))

	handle("POST /api/channels/{channelCode}/messages", xhttp.JSONHandler(api.createMessage))
}
