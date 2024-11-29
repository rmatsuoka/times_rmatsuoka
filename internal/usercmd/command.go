package usercmd

import "github.com/rmatsuoka/times_rmatsuoka/internal/repository"

// Command provides a general manipulation for user repository.
type Command struct {
	Users repository.Users
}

var Default *Command
