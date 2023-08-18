package usecase

import (
	"context"
	"fmt"

	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository `di.inject:"userRepo"`
}

func (u *UserUsecase) FindAllUsers() (*dto.FindAllUsersResponse, error) {
	users, err := u.userRepo.FindAllUsers(context.Background())
	if err != nil {
		return nil, err
	}

	data := dto.FindAllUsersResponseData{}
	for _, usr := range users {
		data = append(data, dto.FindOneUserResponseData{
			ID:        usr.ID,
			FullName:  usr.FullName,
			Username:  usr.Username,
			Email:     usr.Email,
			Role:      usr.Role,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: usr.UpdatedAt,
		})
	}

	response := dto.FindAllUsersResponse{
		Data: data,
	}

	return &response, nil
}

func (u *UserUsecase) FindOneUser(
	userId int32,
) (*dto.FindOneUserResponse, error) {
	usr, err := u.userRepo.FindOneUser(context.Background(), userId)
	if err != nil {
		return nil, err
	}

	response := dto.FindOneUserResponse{
		Data: dto.FindOneUserResponseData{
			ID:        usr.ID,
			FullName:  usr.FullName,
			Username:  usr.Username,
			Email:     usr.Email,
			Role:      usr.Role,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: usr.UpdatedAt,
		},
	}

	return &response, nil
}

func (u *UserUsecase) UpdateUser(
	userId int32,
	payload *dto.UpdateUserRequest,
) (*dto.UpdateUserResponse, error) {
	ctx := context.Background()

	if _, errVerifyUser := u.userRepo.FindOneUser(
		ctx,
		userId,
	); errVerifyUser != nil {
		return nil, errVerifyUser
	}

	updatedUser, errUpdate := u.userRepo.UpdateUser(
		ctx,
		userId, &entity.User{
			FullName: payload.FullName,
			Username: payload.Username,
		})
	if errUpdate != nil {
		return nil, errUpdate
	}

	response := dto.UpdateUserResponse{
		Message: "user data has been updated successfully",
		Data: dto.UpdateUserResponseData{
			ID:        updatedUser.ID,
			FullName:  updatedUser.FullName,
			Username:  updatedUser.Username,
			Email:     updatedUser.Email,
			Role:      updatedUser.Role,
			UpdatedAt: updatedUser.UpdatedAt,
		},
	}

	return &response, nil
}

func (u *UserUsecase) DeleteUser(
	userId int32,
) (*dto.DeleteUserResponse, error) {
	ctx := context.Background()

	if _, err := u.userRepo.FindOneUser(ctx, userId); err != nil {
		return nil, err
	}

	if err := u.userRepo.DeleteUser(ctx, userId); err != nil {
		return nil, err
	}

	response := dto.DeleteUserResponse{
		Message: fmt.Sprintf("user with id %d has been deleted", userId),
	}

	return &response, nil
}
