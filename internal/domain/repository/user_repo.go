package repository

import (
	"context"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
)

type UserRepository interface {
	Register(ctx context.Context, newUser *entity.User) (*entity.User, error)
	FindAllUsers(ctx context.Context) ([]*entity.User, error)
	FindOneUser(ctx context.Context, userId int32) (*entity.User, error)
	FindOneUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(
		ctx context.Context,
		userId int32,
		newUser *entity.User,
	) (*entity.User, error)
	UpdateEmail(
		ctx context.Context,
		userId int32,
		email string,
	) (*entity.User, error)
	UpdatePassword(ctx context.Context, userId int32, password string) error
	DeleteUser(ctx context.Context, userId int32) error
}
