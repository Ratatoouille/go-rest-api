package store

import "restApi/internal/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error

	FindByEmail(string) (*model.User, error)
}
