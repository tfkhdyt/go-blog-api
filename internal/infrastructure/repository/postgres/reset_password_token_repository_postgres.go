package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgtype"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/database/postgres/sqlc"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type ResetPasswordTokenRepositoryPostgres struct {
	db sqlc.Querier `di.inject:"database"`
}

func (r *ResetPasswordTokenRepositoryPostgres) AddToken(
	ctx context.Context,
	userID int32,
	token *entity.ResetPasswordToken,
) (*entity.ResetPasswordToken, error) {
	result, err := r.db.AddResetPasswordToken(
		ctx,
		sqlc.AddResetPasswordTokenParams{
			Token:     token.Token,
			UserID:    pgtype.Int4{Int32: userID, Valid: true},
			ExpiresAt: pgtype.Timestamp{Time: token.ExpiresAt, Valid: true},
		},
	)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, exception.
			NewHTTPError(500, "failed to add reset-password token")
	}

	return &entity.ResetPasswordToken{
		UserID:    result.UserID.Int32,
		Token:     result.Token,
		ExpiresAt: result.ExpiresAt.Time,
	}, nil
}

func (r *ResetPasswordTokenRepositoryPostgres) FindToken(
	ctx context.Context,
	token string,
) (*entity.ResetPasswordToken, error) {
	result, err := r.db.FindResetPasswordToken(ctx, token)
	if err != nil {
		return nil, exception.NewHTTPError(404, "token is not found in database")
	}

	return &entity.ResetPasswordToken{
		Token:     result.Token,
		ExpiresAt: result.ExpiresAt.Time,
		UserID:    result.UserID.Int32,
	}, nil
}

func (r *ResetPasswordTokenRepositoryPostgres) DeleteToken(
	ctx context.Context,
	token string,
) error {
	if err := r.db.DeleteResetPasswordToken(ctx, token); err != nil {
		log.Println("ERROR:", err)
		return exception.NewHTTPError(500, "failed to delete reset-password token")
	}

	return nil
}
