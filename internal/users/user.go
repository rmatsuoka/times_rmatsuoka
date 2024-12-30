package users

import (
	"fmt"
	"regexp"
	"time"
)

type ID string

type User struct {
	ID        ID        `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) UserCode() string { return u.Code }
func (u *User) UserName() string { return u.Name }

type Creating interface {
	UserCode() string
	UserName() string
}

type ValidCreating struct {
	Creating
}

var validCode = regexp.MustCompile("[A-Za-z0-9_]{3,30}")

func ValidateCreating(u Creating) (ValidCreating, error) {
	if !validCode.MatchString(u.UserCode()) {
		return ValidCreating{}, fmt.Errorf("invalid user code %q", u.UserCode())
	}
	if len(u.UserName()) == 0 {
		return ValidCreating{}, fmt.Errorf("invalid user name %q", u.UserName())
	}
	return ValidCreating{u}, nil
}
