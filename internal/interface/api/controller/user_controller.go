package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/application/usecase"
	"codeberg.org/tfkhdyt/blog-api/pkg/auth"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
	"codeberg.org/tfkhdyt/blog-api/pkg/validator"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func (u *UserController) FindAllUsers(c *fiber.Ctx) error {
	users, err := u.userUsecase.FindAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

func (u *UserController) FindOneUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return exception.NewHTTPError(400, "invalid user id")
	}

	usr, errFind := u.userUsecase.FindOneUser(uint(userId))
	if errFind != nil {
		return errFind
	}

	return c.JSON(fiber.Map{
		"data": usr,
	})
}

func (u *UserController) FindMyUser(c *fiber.Ctx) error {
	userId := auth.GetUserIDFromClaims(c)

	usr, errFind := u.userUsecase.FindOneUser(uint(userId))
	if errFind != nil {
		return errFind
	}

	return c.JSON(fiber.Map{
		"data": usr,
	})
}

func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return exception.NewHTTPError(400, "invalid user id")
	}

	payload := new(dto.UpdateUserRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return validator.NewValidationError(err)
	}

	updatedUser, errUpdate := u.userUsecase.UpdateUser(uint(userId), payload)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(fiber.Map{
		"data": updatedUser,
	})
}

func (u *UserController) UpdateMyUser(c *fiber.Ctx) error {
	userId := auth.GetUserIDFromClaims(c)

	payload := new(dto.UpdateUserRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return validator.NewValidationError(err)
	}

	updatedUser, errUpdate := u.userUsecase.UpdateUser(uint(userId), payload)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(fiber.Map{
		"data": updatedUser,
	})
}

func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return exception.NewHTTPError(400, "invalid user id")
	}

	response, errDelete := u.userUsecase.DeleteUser(uint(userId))
	if errDelete != nil {
		return errDelete
	}

	return c.JSON(response)
}

func (u *UserController) DeleteMyUser(c *fiber.Ctx) error {
	userId := auth.GetUserIDFromClaims(c)

	response, errDelete := u.userUsecase.DeleteUser(uint(userId))
	if errDelete != nil {
		return errDelete
	}

	return c.JSON(response)
}
