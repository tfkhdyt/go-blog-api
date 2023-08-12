package repository

import "codeberg.org/tfkhdyt/blog-api/internal/domain/entity"

type AuthRepository interface {
	AddToken(auth *entity.Auth) (*entity.Auth, error)
	VerifyToken(token string) error
	RemoveToken(token string) error
}
