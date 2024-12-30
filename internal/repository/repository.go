package repository

import "sync"

type Repository struct {
	Users    Users
	Channels Channels
	Messages Messages
}

var defaultRepository *Repository

func Default() *Repository {
	return defaultRepository
}

var setDefaultRepositoryMu sync.Mutex

func InitDefault(repo *Repository) {
	setDefaultRepositoryMu.Lock()
	defer setDefaultRepositoryMu.Unlock()

	if defaultRepository != nil {
		panic("set default repository twice time")
	}

	defaultRepository = repo
}
