package usercmd

import "github.com/rmatsuoka/times_rmatsuoka/internal/repository"

// Command provides a general manipulation for user repository.
type Command struct {
	Repository *repository.Repository
}

var Default *Command
