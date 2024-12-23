package infratypes

import (
	"strconv"

	"github.com/rmatsuoka/times_rmatsuoka/internal/infra/schema"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
)

type UserID int64

var _ users.ID = UserID(0)

func (i UserID) String() string { return strconv.FormatInt(int64(i), 10) }
func (i UserID) UserID()        {}

func UsersUser(su *schema.User) *users.User {
	return &users.User{
		ID:        UserID(su.ID),
		Code:      su.Code,
		Name:      su.Name,
		CreatedAt: su.CreatedAt,
		UpdatedAt: su.UpdatedAt,
	}
}
