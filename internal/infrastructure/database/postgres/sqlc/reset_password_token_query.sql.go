// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: reset_password_token_query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addResetPasswordToken = `-- name: AddResetPasswordToken :one
INSERT INTO reset_password_token (
  token, user_id, expires_at
) VALUES (
  $1, $2, $3
) RETURNING token, expires_at, user_id
`

type AddResetPasswordTokenParams struct {
	Token     string
	UserID    pgtype.Int4
	ExpiresAt pgtype.Timestamp
}

func (q *Queries) AddResetPasswordToken(ctx context.Context, arg AddResetPasswordTokenParams) (ResetPasswordToken, error) {
	row := q.db.QueryRow(ctx, addResetPasswordToken, arg.Token, arg.UserID, arg.ExpiresAt)
	var i ResetPasswordToken
	err := row.Scan(&i.Token, &i.ExpiresAt, &i.UserID)
	return i, err
}

const deleteResetPasswordToken = `-- name: DeleteResetPasswordToken :exec
DELETE FROM reset_password_token
WHERE token = $1
`

func (q *Queries) DeleteResetPasswordToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteResetPasswordToken, token)
	return err
}

const findResetPasswordToken = `-- name: FindResetPasswordToken :one
SELECT token, expires_at, user_id FROM reset_password_token 
WHERE token = $1
`

func (q *Queries) FindResetPasswordToken(ctx context.Context, token string) (ResetPasswordToken, error) {
	row := q.db.QueryRow(ctx, findResetPasswordToken, token)
	var i ResetPasswordToken
	err := row.Scan(&i.Token, &i.ExpiresAt, &i.UserID)
	return i, err
}
