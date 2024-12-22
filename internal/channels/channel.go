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

type MemberID struct {
	Channel ID
	User    users.ID
	Role    Role
}

type Member struct {
	Channel ID
	*users.User
	Role Role
}

func MemberFromID(user *users.User, mid *MemberID) *Member {
	return &Member{
		Channel: mid.Channel,
		User:    user,
		Role:    mid.Role,
	}
}

type Role int

const (
	RoleOwenr Role = 1 + iota
	RoleAdmin
	RoleWriter
	RoleReader
)
