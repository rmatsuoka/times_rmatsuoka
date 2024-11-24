package xhttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func JSONHandler[T any](f JSONHandlerFunc[T]) http.Handler {
	return f
}

type JSONHandlerFunc[T any] func(w http.ResponseWriter, req *http.Request, body T)

type statusJSON struct {
	Message string `json:"message,omitempty"`
}

func (h JSONHandlerFunc[T]) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rc := http.MaxBytesReader(w, req.Body, 1<<19)
	b, err := io.ReadAll(rc)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, statusJSON{
			Message: err.Error(),
		})
		return
	}
	var body T
	if err = json.Unmarshal(b, &body); err != nil {
		WriteJSON(w, http.StatusBadRequest, statusJSON{
			Message: err.Error(),
		})
		return
	}
	h(w, req, body)
}

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	buf, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(buf)
}
