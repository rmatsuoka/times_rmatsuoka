package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/rmatsuoka/times_rmatsuoka/internal/api"
	"github.com/rmatsuoka/times_rmatsuoka/internal/infra"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xhttp"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xslog"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

func main() {
	logger := slog.New(xslog.NewContextHandler(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	})))
	slog.SetDefault(logger)

	infra.Init()
	db, err := sql.Open("sqlite3", "file:local.db?_fk=1")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	a := &api.API{
		DB: &xsql.SQLDB{DB: db},
	}
	a.Install(mux.Handle)

	addr := ":8000"
	log.Printf("listen on %s", addr)
	handler := xhttp.LoggingHandler(mux)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
