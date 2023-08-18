package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/database/postgres/sqlc"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type UserRepositoryPostgres struct {
	db sqlc.Querier `di.inject:"database"`
}

func (u *UserRepositoryPostgres) Register(
	ctx context.Context,
	newUser *entity.User,
) (*entity.User, error) {
	result, err := u.db.CreateUser(ctx, sqlc.CreateUserParams{
		FullName: newUser.FullName,
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
		Role:     sqlc.NullRole{Role: sqlc.Role(newUser.Role), Valid: true},
	})
	if err != nil {
		log.Println("ERROR:", err)
		return nil, exception.NewHTTPError(500, "failed to register new user")
	}

	return &entity.User{
		CreatedAt: result.CreatedAt.Time,
		FullName:  result.FullName,
		Username:  result.Username,
		Email:     result.Email,
		Role:      entity.Role(result.Role.Role),
		ID:        result.ID,
	}, nil
}

func (u *UserRepositoryPostgres) FindAllUsers(
	ctx context.Context,
) ([]*entity.User, error) {
	result, err := u.db.FindAllUsers(ctx)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, exception.NewHTTPError(500, "failed to find all users")
	}

	users := []*entity.User{}
	for _, res := range result {
		users = append(users, &entity.User{
			CreatedAt: res.CreatedAt.Time,
			UpdatedAt: res.UpdatedAt.Time,
			FullName:  res.FullName,
			Username:  res.Username,
			Email:     res.Email,
			Role:      entity.Role(res.Role.Role),
			ID:        res.ID,
		})
	}

	return users, nil
}

func (u *UserRepositoryPostgres) FindOneUser(
	ctx context.Context,
	userId int32,
) (*entity.User, error) {
	result, err := u.db.FindOneUserByID(ctx, userId)
	if err != nil {
		return nil, exception.
			NewHTTPError(404, fmt.Sprintf("user with id %d is not found", userId))
	}

	return &entity.User{
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
		FullName:  result.FullName,
		Username:  result.Username,
		Email:     result.Email,
		Password:  result.Password,
		Role:      entity.Role(result.Role.Role),
		ID:        result.ID,
	}, nil
}

func (u *UserRepositoryPostgres) FindOneUserByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {
	result, err := u.db.FindOneUserByEmail(ctx, email)
	if err != nil {
		return nil, exception.
			NewHTTPError(404, fmt.Sprintf("user with email %s is not found", email))
	}

	return &entity.User{
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
		FullName:  result.FullName,
		Username:  result.Username,
		Email:     result.Email,
		Password:  result.Password,
		Role:      entity.Role(result.Role.Role),
		ID:        result.ID,
	}, nil
}

func (u *UserRepositoryPostgres) UpdateUser(
	ctx context.Context,
	userId int32,
	newUser *entity.User,
) (*entity.User, error) {
	result, err := u.db.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:        userId,
		FullName:  newUser.FullName,
		Username:  newUser.Username,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, exception.NewHTTPError(
			500,
			fmt.Sprintf("failed to update user with id %d", userId),
		)
	}

	return &entity.User{
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
		FullName:  result.FullName,
		Username:  result.Username,
		Email:     result.Email,
		Role:      entity.Role(result.Role.Role),
		ID:        result.ID,
	}, nil
}

func (u *UserRepositoryPostgres) UpdateEmail(
	ctx context.Context,
	userId int32,
	email string,
) (*entity.User, error) {
	result, err := u.db.UpdateEmail(ctx, sqlc.UpdateEmailParams{
		ID:        userId,
		Email:     email,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, exception.NewHTTPError(
			500,
			fmt.Sprintf("failed to update user with id %d", userId),
		)
	}

	return &entity.User{
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
		FullName:  result.FullName,
		Username:  result.Username,
		Email:     result.Email,
		Role:      entity.Role(result.Role.Role),
		ID:        result.ID,
	}, nil
}

func (u *UserRepositoryPostgres) UpdatePassword(
	ctx context.Context,
	userId int32,
	password string,
) error {
	if err := u.db.UpdatePassword(
		ctx,
		sqlc.UpdatePasswordParams{
			ID:        userId,
			Password:  password,
			UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		},
	); err != nil {
		return exception.NewHTTPError(
			500,
			fmt.Sprintf("failed to update user with id %d", userId),
		)
	}

	return nil
}

func (u *UserRepositoryPostgres) DeleteUser(
	ctx context.Context,
	userId int32,
) error {
	if err := u.db.DeleteUser(ctx, userId); err != nil {
		return exception.NewHTTPError(
			500,
			fmt.Sprintf("failed to delete user with id %d", userId),
		)
	}

	return nil
}
