package teststore

import (
	"restApi/internal/model"
	"restApi/internal/store"
)

// Store
type Store struct {
	userRepo *UserRepository
}

// New
func New() *Store {
	return &Store{}
}

// User
func (s *Store) User() store.UserRepository {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepo
}
