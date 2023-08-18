package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgtype"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/database/postgres/sqlc"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type ChangeEmailTokenRepositoryPostgres struct {
	db sqlc.Querier `di.inject:"database"`
}

func (r *ChangeEmailTokenRepositoryPostgres) AddToken(
	ctx context.Context,
	userId int32,
	request *entity.ChangeEmailRequest,
) (*entity.ChangeEmailRequest, error) {
	result, err := r.db.AddChangeEmailToken(
		ctx,
		sqlc.AddChangeEmailTokenParams{
			Token:     request.Token,
			NewEmail:  request.NewEmail,
			UserID:    pgtype.Int4{Int32: userId, Valid: true},
			ExpiresAt: pgtype.Timestamp{Time: request.ExpiresAt, Valid: true},
		},
	)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, exception.
			NewHTTPError(500, "failed to add change-email request")
	}

	return &entity.ChangeEmailRequest{
		Token:     result.Token,
		NewEmail:  result.NewEmail,
		UserID:    result.UserID.Int32,
		ExpiresAt: result.ExpiresAt.Time,
	}, nil
}

func (r *ChangeEmailTokenRepositoryPostgres) FindToken(
	ctx context.Context,
	token string,
) (*entity.ChangeEmailRequest, error) {
	result, err := r.db.FindChangeEmailToken(ctx, token)
	if err != nil {
		return nil, exception.NewHTTPError(404, "token is not found in database")
	}

	return &entity.ChangeEmailRequest{
		Token:     result.Token,
		NewEmail:  result.NewEmail,
		UserID:    result.UserID.Int32,
		ExpiresAt: result.ExpiresAt.Time,
	}, nil
}

func (r *ChangeEmailTokenRepositoryPostgres) DeleteToken(
	ctx context.Context,
	token string,
) error {
	if err := r.db.DeleteChangeEmailToken(
		ctx,
		token,
	); err != nil {
		return exception.NewHTTPError(500, "failed to delete change-email token")
	}

	return nil
}
