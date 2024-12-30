package api

import (
	"errors"
	"net/http"

	"github.com/rmatsuoka/times_rmatsuoka/internal/types"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xhttp"
)

type status string

const (
	statusError status = "error"
	statusOK    status = "ok"
)

type statusJSON struct {
	Status  status `json:"status"`
	Message string `json:"message,omitempty"`
}

var errStatus = map[error]int{
	types.ErrExist: http.StatusConflict,
}

var empty = struct{}{}

func writeResult(w http.ResponseWriter, data any, err error) {
	if err != nil {
		for target, status := range errStatus {
			if errors.Is(err, target) {
				xhttp.WriteJSON(w, status, statusJSON{
					Status:  statusError,
					Message: err.Error(),
				})
				return
			}
		}

		xhttp.WriteJSON(w, http.StatusInternalServerError, statusJSON{
			Status:  statusError,
			Message: err.Error(),
		})
		return
	}

	if data == empty {
		xhttp.WriteJSON(w, http.StatusOK, statusJSON{
			Status: statusOK,
		})
		return
	}

	xhttp.WriteJSON(w, http.StatusOK, data)
}
