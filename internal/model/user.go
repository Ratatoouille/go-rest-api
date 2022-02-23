package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// User
type User struct {
	ID                int
	Email             string `validate:"required,email"`
	Password          string `validate:"required,min=8"`
	EncryptedPassword string
}

// Validate
func (u *User) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}

// BeforeCreate
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

// encryptString
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
