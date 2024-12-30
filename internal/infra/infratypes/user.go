package infratypes

import (
	"strconv"

	"github.com/rmatsuoka/times_rmatsuoka/internal/infra/schema"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
)

func UserID(i int64) users.ID {
	return users.ID(strconv.FormatInt(i, 10))
}

func UsersUser(su *schema.User) *users.User {
	return &users.User{
		ID:        UserID(su.ID),
		Code:      su.Code,
		Name:      su.Name,
		CreatedAt: su.CreatedAt,
		UpdatedAt: su.UpdatedAt,
	}
}
