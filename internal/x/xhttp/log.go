package xhttp

import (
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
			"pattern", req.Pattern,
			"remoteAddr", req.RemoteAddr,
			"userAgent", req.UserAgent(),
		))
		logw := &logResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(logw, req.WithContext(ctx))
		slog.InfoContext(ctx, "request", "statusCode", logw.statusCode)
	})
}

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *logResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *logResponseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

var _ http.ResponseWriter = (*logResponseWriter)(nil)
