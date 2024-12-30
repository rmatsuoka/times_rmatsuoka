package infra

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/rmatsuoka/times_rmatsuoka/internal/infra/userinfra"
	"github.com/rmatsuoka/times_rmatsuoka/internal/repository"
)

func Init() {
	repository.InitDefault(&repository.Repository{
		Users:    userinfra.Users{},
		Channels: nil,
		Messages: nil,
	})
}
