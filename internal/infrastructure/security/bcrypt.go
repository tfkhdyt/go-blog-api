package security

import (
	"golang.org/x/crypto/bcrypt"

	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type PasswordHashService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) error
}

type BcryptService struct{}

func (b *BcryptService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", exception.NewHTTPError(500, "failed to hash password")
	}

	return string(hashedPassword), nil
}

func (b *BcryptService) ComparePassword(hashedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return exception.NewHTTPError(400, "invalid password")
	}

	return nil
}
