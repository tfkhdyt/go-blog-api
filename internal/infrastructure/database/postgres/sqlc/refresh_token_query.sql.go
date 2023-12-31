// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: refresh_token_query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addRefreshToken = `-- name: AddRefreshToken :one
INSERT INTO refresh_token (
  token, user_id
) VALUES (
  $1, $2
) RETURNING token, user_id
`

type AddRefreshTokenParams struct {
	Token  string
	UserID pgtype.Int4
}

func (q *Queries) AddRefreshToken(ctx context.Context, arg AddRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, addRefreshToken, arg.Token, arg.UserID)
	var i RefreshToken
	err := row.Scan(&i.Token, &i.UserID)
	return i, err
}

const deleteRefreshToken = `-- name: DeleteRefreshToken :exec
DELETE FROM refresh_token
WHERE token = $1
`

func (q *Queries) DeleteRefreshToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteRefreshToken, token)
	return err
}

const findRefreshToken = `-- name: FindRefreshToken :one
SELECT token, user_id FROM refresh_token 
WHERE token = $1
`

func (q *Queries) FindRefreshToken(ctx context.Context, token string) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, findRefreshToken, token)
	var i RefreshToken
	err := row.Scan(&i.Token, &i.UserID)
	return i, err
}
