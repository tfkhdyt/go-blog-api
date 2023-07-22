package user

import (
	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *userHandler {
	return &userHandler{userService}
}

func (u *userHandler) Register(c *fiber.Ctx) error {
	payload := new(user.RegisterRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "failed to parse body")
	}

	registeredUser, err := u.userService.Register(payload)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": registeredUser,
	})
}

func (u *userHandler) FindAllUsers(c *fiber.Ctx) error {
	users, err := u.userService.FindAllUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

func (u *userHandler) FindOneUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	usr, errFind := u.userService.FindOneUser(uint(userId))
	if errFind != nil {
		return errFind
	}

	return c.JSON(fiber.Map{
		"data": usr,
	})
}

func (u *userHandler) UpdateCake(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	payload := new(user.UpdateUserRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "failed to parse body")
	}

	updatedUser, errUpdate := u.userService.UpdateUser(uint(userId), payload)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(fiber.Map{
		"data": updatedUser,
	})
}

func (u *userHandler) DeleteCake(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	response, errDelete := u.userService.DeleteUser(uint(userId))
	if errDelete != nil {
		return errDelete
	}

	return c.JSON(response)
}
