package channelcmd

import (
	"cmp"

	"github.com/rmatsuoka/times_rmatsuoka/internal/repository"
)

type Command struct {
	Repositories *repository.Repositories
}

func (c *Command) repositories() *repository.Repositories {
	return cmp.Or(c.Repositories, repository.Default())
}

var Default = &Command{}
