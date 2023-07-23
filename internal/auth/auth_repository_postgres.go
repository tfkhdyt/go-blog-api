package auth

import (
	"fmt"

	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/auth"
)

type authRepoPostgres struct {
	db *gorm.DB
}

func NewAuthRepoPostgres(db *gorm.DB) *authRepoPostgres {
	return &authRepoPostgres{db}
}

func (a *authRepoPostgres) AddToken(auth *auth.Auth) (*auth.Auth, error) {
	if err := a.db.Create(auth).Error; err != nil {
		return nil, fmt.Errorf("failed to add refresh token")
	}

	return auth, nil
}

func (a *authRepoPostgres) VerifyToken(token string) error {
	tkn := new(auth.Auth)
	if err := a.db.First(tkn, "refresh_token = ?", token).Error; err != nil {
		return fmt.Errorf("token is not found")
	}

	return nil
}

func (a *authRepoPostgres) RemoveToken(token string) error {
	if err := a.db.Delete(&auth.Auth{}, "refresh_token = ?", token).Error; err != nil {
		return fmt.Errorf("failed to delete refresh token")
	}

	return nil
}
