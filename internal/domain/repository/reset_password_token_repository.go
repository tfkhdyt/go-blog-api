package repository

import (
	"context"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
)

type ResetPasswordTokenRepository interface {
	AddToken(
		ctx context.Context,
		userID int32,
		token *entity.ResetPasswordToken,
	) (*entity.ResetPasswordToken, error)
	FindToken(
		ctx context.Context,
		token string,
	) (*entity.ResetPasswordToken, error)
	DeleteToken(ctx context.Context, token string) error
}
