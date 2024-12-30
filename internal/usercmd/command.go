package usercmd

import (
	"cmp"

	"github.com/rmatsuoka/times_rmatsuoka/internal/repository"
)

// Command provides a general manipulation for user repository.
type Command struct {
	Repositories *repository.Repositories
}

func (c *Command) repositories() *repository.Repositories {
	return cmp.Or(c.Repositories, repository.Default())
}

func (c *Command) Users() repository.Users {
	return c.repositories().Users
}

var Default = &Command{}
