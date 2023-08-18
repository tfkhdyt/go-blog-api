package repository

import (
	"context"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
)

type AuthRepository interface {
	AddToken(
		ctx context.Context,
		userId int32,
		refreshToken *entity.RefreshToken,
	) (*entity.RefreshToken, error)
	VerifyToken(ctx context.Context, token string) error
	DeleteToken(ctx context.Context, token string) error
}
