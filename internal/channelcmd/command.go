package channelcmd

import "github.com/rmatsuoka/times_rmatsuoka/internal/repository"

type Command struct {
	Repository *repository.Repository
}

var Default *Command
