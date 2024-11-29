package api

import (
	"net/http"

	"github.com/rmatsuoka/times_rmatsuoka/internal/channelcmd"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xhttp"
)

type createMessageBody struct {
	Text string `json:"text"`
}

func (api *API) createMessage(w http.ResponseWriter, req *http.Request, body *createMessageBody) {
	channel := req.PathValue("channelCode")
	id, err := channelcmd.CreateMessage(req.Context(), api.DB, channel, &channelcmd.CreatingMessage{
		Text: body.Text,
	})
	if err != nil {
		xhttp.WriteJSON(w, 500, map[string]any{"message": err.Error()})
		return
	}
	xhttp.WriteJSON(w, 200, map[string]any{"id": id})
}
