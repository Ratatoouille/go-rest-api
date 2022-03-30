package sqlstore

import (
	"restApi/internal/store"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Store
type Store struct {
	db       *pgxpool.Pool
	userRepo *UserRepository
}

// New
func New(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

// User
func (s *Store) User() store.UserRepository {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepository{
		store: s,
	}

	return s.userRepo
}
