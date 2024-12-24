package testdb

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"

	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

var open = sync.OnceValues(func() (*sql.DB, error) {
	return sql.Open("sqlite3", "file:../../../local.db?_fk=1")
})

func Open() xsql.DB {
	db, err := open()
	if err != nil {
		panic(err)
	}
	return &xsql.SQLDB{DB: db}
}
