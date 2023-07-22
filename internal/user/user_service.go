package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

type userService struct {
	userRepo user.UserRepository
}

func NewUserService(userRepo user.UserRepository) *userService {
	return &userService{userRepo}
}

func (u *userService) Register(payload *user.RegisterRequest) (*user.RegisterResponse, error) {
	newUser, err := payload.Validate()
	if err != nil {
		return nil, err
	}

	if err := newUser.HashPassword(); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	registeredUser, errRegister := u.userRepo.Register(newUser)
	if errRegister != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errRegister.Error())
	}

	response := user.RegisterResponse{
		ID:        registeredUser.ID,
		FullName:  registeredUser.FullName,
		Username:  registeredUser.Username,
		Email:     registeredUser.Email,
		Role:      registeredUser.Role,
		CreatedAt: registeredUser.CreatedAt,
	}

	return &response, nil
}

func (u *userService) FindAllUsers() (*user.FindAllUsersResponse, error) {
	users, err := u.userRepo.FindAllUsers()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := user.FindAllUsersResponse{}
	for _, usr := range *users {
		response = append(response, user.FindOneUserResponse{
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

func (u *userService) FindOneUser(userId uint) (*user.FindOneUserResponse, error) {
	usr, err := u.userRepo.FindOneUser(userId)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := user.FindOneUserResponse{
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

func (u *userService) UpdateUser(
	userId uint,
	payload *user.UpdateUserRequest,
) (*user.UpdateUserResponse, error) {
	newUser, err := payload.Validate()
	if err != nil {
		return nil, err
	}

	oldUser, errVerifyUser := u.userRepo.FindOneUser(userId)
	if errVerifyUser != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, errVerifyUser.Error())
	}

	updatedUser, errUpdate := u.userRepo.UpdateUser(oldUser, newUser)
	if errUpdate != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errUpdate.Error())
	}

	response := user.UpdateUserResponse{
		ID:        updatedUser.ID,
		FullName:  updatedUser.FullName,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		Role:      updatedUser.Role,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return &response, nil
}

func (u *userService) DeleteUser(userId uint) (*user.DeleteUserResponse, error) {
	if _, err := u.userRepo.FindOneUser(userId); err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if err := u.userRepo.DeleteUser(userId); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := user.DeleteUserResponse{
		Message: fmt.Sprintf("user with id %d has been deleted", userId),
	}

	return &response, nil
}
