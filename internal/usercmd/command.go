package usercmd

import (
	"cmp"

	"github.com/rmatsuoka/times_rmatsuoka/internal/repository"
)

// Command provides a general manipulation for user repository.
type Command struct {
	Repository *repository.Repository
}

func (c *Command) repository() *repository.Repository {
	return cmp.Or(c.Repository, repository.Default())
}

var Default = &Command{}
