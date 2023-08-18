package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgtype"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/database/postgres/sqlc"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type AuthRepositoryPostgres struct {
	db sqlc.Querier `di.inject:"database"`
}

func (a *AuthRepositoryPostgres) AddToken(
	ctx context.Context,
	userId int32,
	refreshToken *entity.RefreshToken,
) (*entity.RefreshToken, error) {
	result, err := a.db.AddRefreshToken(ctx, sqlc.AddRefreshTokenParams{
		Token:  refreshToken.Token,
		UserID: pgtype.Int4{Int32: userId, Valid: true},
	})
	if err != nil {
		log.Println("ERROR:", err)
		return nil, exception.NewHTTPError(500, "failed to add refresh token")
	}

	return &entity.RefreshToken{
		Token:  result.Token,
		UserID: result.UserID.Int32,
	}, nil
}

func (a *AuthRepositoryPostgres) VerifyToken(
	ctx context.Context,
	token string,
) error {
	if _, err := a.db.FindRefreshToken(ctx, token); err != nil {
		return exception.
			NewHTTPError(401, "refresh token is not found in database")
	}

	return nil
}

func (a *AuthRepositoryPostgres) DeleteToken(
	ctx context.Context,
	token string,
) error {
	if err := a.db.DeleteRefreshToken(ctx, token); err != nil {
		log.Println("ERROR:", err)
		return exception.NewHTTPError(500, "failed to delete refresh token")
	}

	return nil
}
