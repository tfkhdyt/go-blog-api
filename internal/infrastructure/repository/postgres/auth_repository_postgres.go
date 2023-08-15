package postgres

import (
	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type AuthRepositoryPostgres struct {
	db *gorm.DB `di.inject:"database"`
}

func (a *AuthRepositoryPostgres) AddToken(
	auth *entity.Auth,
) (*entity.Auth, error) {
	if err := a.db.Create(auth).Error; err != nil {
		return nil, exception.NewHTTPError(500, "failed to add refresh token")
	}

	return auth, nil
}

func (a *AuthRepositoryPostgres) VerifyToken(token string) error {
	tkn := new(entity.Auth)
	if err := a.db.First(tkn, "refresh_token = ?", token).Error; err != nil {
		return exception.NewHTTPError(401, "token is not found")
	}

	return nil
}

func (a *AuthRepositoryPostgres) RemoveToken(token string) error {
	if err := a.db.Delete(
		&entity.Auth{},
		"refresh_token = ?",
		token,
	).Error; err != nil {
		return exception.NewHTTPError(500, "failed to delete refresh token")
	}

	return nil
}
