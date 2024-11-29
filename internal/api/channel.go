package api

import (
	"net/http"

	"github.com/rmatsuoka/times_rmatsuoka/internal/channelcmd"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xhttp"
)

func (api *API) createChannel(w http.ResponseWriter, req *http.Request, channel *channelcmd.Creating) {
	id, err := channelcmd.Create(req.Context(), api.DB, channel)
	if err != nil {
		xhttp.WriteJSON(w, 500, map[string]any{"message": err.Error()})
		return
	}
	xhttp.WriteJSON(w, 200, map[string]any{"id": id})
}
