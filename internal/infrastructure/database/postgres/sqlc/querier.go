// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package sqlc

import (
	"context"
)

type Querier interface {
	AddChangeEmailToken(ctx context.Context, arg AddChangeEmailTokenParams) (ChangeEmailToken, error)
	AddRefreshToken(ctx context.Context, arg AddRefreshTokenParams) (RefreshToken, error)
	AddResetPasswordToken(ctx context.Context, arg AddResetPasswordTokenParams) (ResetPasswordToken, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteChangeEmailToken(ctx context.Context, token string) error
	DeleteRefreshToken(ctx context.Context, token string) error
	DeleteResetPasswordToken(ctx context.Context, token string) error
	DeleteUser(ctx context.Context, id int32) error
	FindAdmin(ctx context.Context) ([]User, error)
	FindAllUsers(ctx context.Context) ([]FindAllUsersRow, error)
	FindChangeEmailToken(ctx context.Context, token string) (ChangeEmailToken, error)
	FindOneUserByEmail(ctx context.Context, email string) (User, error)
	FindOneUserByID(ctx context.Context, id int32) (User, error)
	FindRefreshToken(ctx context.Context, token string) (RefreshToken, error)
	FindResetPasswordToken(ctx context.Context, token string) (ResetPasswordToken, error)
	UpdateEmail(ctx context.Context, arg UpdateEmailParams) (UpdateEmailRow, error)
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
}

var _ Querier = (*Queries)(nil)
