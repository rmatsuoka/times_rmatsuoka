package channels

import (
	"fmt"
	"regexp"
	"time"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
)

type ID interface {
	String() string
	ChannelID()
}

type Channel struct {
	ID        ID
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Creating interface {
	ChannelCode() string
}

type ValidCreating struct {
	Creating
}

var validCode = regexp.MustCompile("[A-Za-z_]{,30}")

func ValidateCreating(c Creating) (ValidCreating, error) {
	if !validCode.MatchString(c.ChannelCode()) {
		return ValidCreating{}, fmt.Errorf("invalid code")
	}
	return ValidCreating{c}, nil
}

type ChannelUserID struct {
	Channel ID
	User    users.ID
	Role    Role
}

type ChannelUser struct {
	Channel ID
	*users.User
	Role Role
}

func ChannelUserFromID(user *users.User, cuid *ChannelUserID) *ChannelUser {
	return &ChannelUser{
		Channel: cuid.Channel,
		User:    user,
		Role:    cuid.Role,
	}
}

type Role int

const (
	RoleOwenr Role = 1 + iota
	RoleAdmin
	RoleWriter
	RoleReader
)
