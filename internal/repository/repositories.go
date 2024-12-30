package repository

import "sync"

type Repositories struct {
	Users    Users
	Channels Channels
	Messages Messages
}

var defaultRepositories *Repositories

func Default() *Repositories {
	return defaultRepositories
}

var setDefaultRepositoriesMu sync.Mutex

func InitDefault(repo *Repositories) {
	setDefaultRepositoriesMu.Lock()
	defer setDefaultRepositoriesMu.Unlock()

	if defaultRepositories != nil {
		panic("set default repositories twice time")
	}

	defaultRepositories = repo
}
