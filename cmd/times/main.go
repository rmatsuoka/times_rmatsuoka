package times

import (
	"log/slog"
	"net/http"

	"github.com/rmatsuoka/times_rmatsuoka/internal/api"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xhttp"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xslog"
)

func main() {
	logger := slog.New(xslog.NewContextHandler(slog.Default().Handler()))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	a := &api.API{}
	a.Install(mux.Handle)

	handler := xhttp.LoggingHandler(mux)
	http.ListenAndServe(":8080", handler)
}
