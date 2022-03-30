package sqlstore

import (
	"context"
	"restApi/internal/model"
	"restApi/internal/store"

	"github.com/jackc/pgx/v4"
)

// UserRepository
type UserRepository struct {
	store *Store
}

// Create
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	if err := r.store.db.QueryRow(context.Background(),
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		if err == pgx.ErrNoRows {
			return store.ErrRecordNotFound
		}

		return err
	}

	return nil
}

// FindByEmail
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		context.Background(),
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}