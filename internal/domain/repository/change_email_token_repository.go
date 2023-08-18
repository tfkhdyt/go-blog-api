package repository

import (
	"context"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
)

type ChangeEmailTokenRepository interface {
	AddToken(
		ctx context.Context,
		userId int32,
		request *entity.ChangeEmailRequest,
	) (*entity.ChangeEmailRequest, error)
	FindToken(
		ctx context.Context,
		token string,
	) (*entity.ChangeEmailRequest, error)
	DeleteToken(ctx context.Context, token string) error
}
