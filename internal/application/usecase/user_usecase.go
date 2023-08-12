package usecase

import (
	"fmt"

	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository `di.inject:"userRepo"`
}

func (u *UserUsecase) FindAllUsers() (*dto.FindAllUsersResponse, error) {
	users, err := u.userRepo.FindAllUsers()
	if err != nil {
		return nil, err
	}

	response := dto.FindAllUsersResponse{}
	for _, usr := range *users {
		response = append(response, dto.FindOneUserResponse{
			ID:        usr.ID,
			FullName:  usr.FullName,
			Username:  usr.Username,
			Email:     usr.Email,
			Role:      usr.Role,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: usr.UpdatedAt,
		})
	}

	return &response, nil
}

func (u *UserUsecase) FindOneUser(userId uint) (*dto.FindOneUserResponse, error) {
	usr, err := u.userRepo.FindOneUser(userId)
	if err != nil {
		return nil, err
	}

	response := dto.FindOneUserResponse{
		ID:        usr.ID,
		FullName:  usr.FullName,
		Username:  usr.Username,
		Email:     usr.Email,
		Role:      usr.Role,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}

	return &response, nil
}

func (u *UserUsecase) UpdateUser(
	userId uint,
	payload *dto.UpdateUserRequest,
) (*dto.UpdateUserResponse, error) {
	oldUser, errVerifyUser := u.userRepo.FindOneUser(userId)
	if errVerifyUser != nil {
		return nil, errVerifyUser
	}

	updatedUser, errUpdate := u.userRepo.UpdateUser(oldUser, &entity.User{
		FullName: payload.FullName,
		Username: payload.Username,
		Email:    payload.Email,
	})
	if errUpdate != nil {
		return nil, errUpdate
	}

	response := dto.UpdateUserResponse{
		ID:        updatedUser.ID,
		FullName:  updatedUser.FullName,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		Role:      updatedUser.Role,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return &response, nil
}

func (u *UserUsecase) DeleteUser(userId uint) (*dto.DeleteUserResponse, error) {
	if _, err := u.userRepo.FindOneUser(userId); err != nil {
		return nil, err
	}

	if err := u.userRepo.DeleteUser(userId); err != nil {
		return nil, err
	}

	response := dto.DeleteUserResponse{
		Message: fmt.Sprintf("user with id %d has been deleted", userId),
	}

	return &response, nil
}
