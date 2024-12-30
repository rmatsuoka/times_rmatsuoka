package xhttp

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"

	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xslog"
)

var requestAttrKey = xslog.NewAttrKey()

func LoggingHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), requestAttrKey, slog.Group("request",
			"method", req.Method,
			"url", req.URL.String(),
			"remoteAddr", req.RemoteAddr,
			"userAgent", req.UserAgent(),
		))
		logw := &logResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(logw, req.WithContext(ctx))
		slog.InfoContext(ctx, "request",
			"statusCode", logw.statusCode,
			"errorResponseBody", logw.errResponse.String(),
		)
	})
}

type logResponseWriter struct {
	http.ResponseWriter
	statusCode  int
	errResponse *bytes.Buffer
}

func (w *logResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *logResponseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

func (w *logResponseWriter) Write(p []byte) (int, error) {
	if w.errResponse == nil {
		w.errResponse = new(bytes.Buffer)
	}
	if w.statusCode >= 400 {
		w.errResponse.Write(p)
	}
	return w.ResponseWriter.Write(p)
}

var _ http.ResponseWriter = (*logResponseWriter)(nil)
