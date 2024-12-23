package userinfra

import (
	"github.com/rmatsuoka/times_rmatsuoka/internal/repository"
)

type Users struct{}

var _ repository.Users = Users{}
