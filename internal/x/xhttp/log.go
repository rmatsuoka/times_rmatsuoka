package xhttp

import (
	"cmp"
	"log/slog"
	"net/http"
)

func LoggingHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logw := &logResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(logw, req)
		slog.InfoContext(req.Context(), "request",
			"method", req.Method,
			"url", req.URL.String(),
			"pattern", req.Pattern,
			"statusCode", cmp.Or(logw.statusCode, 200),
			"remoteAddr", req.RemoteAddr,
			"userAgent", req.UserAgent(),
		)
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
