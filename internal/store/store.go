package store

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Store
type Store struct {
	config   *Config
	db       *pgxpool.Pool
	userRepo *UserRepository
}

// New
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open
func (s *Store) Open() error {
	db, err := pgxpool.Connect(context.Background(), s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close
func (s *Store) Close() {
	s.db.Close()
}

// User
func (s *Store) User() *UserRepository {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepository{
		store: s,
	}

	return s.userRepo
}
