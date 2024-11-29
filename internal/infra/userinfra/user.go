package userinfra

import (
	"strconv"

	"github.com/rmatsuoka/times_rmatsuoka/internal/repository"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
)

type User struct{}

var _ repository.Users = User{}

type ID int64

var _ users.ID = ID(0)

func (i ID) String() string { return strconv.FormatInt(int64(i), 10) }
func (i ID) UserID()        {}
