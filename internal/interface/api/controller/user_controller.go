package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/application/usecase"
	"codeberg.org/tfkhdyt/blog-api/pkg/auth"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type UserController struct {
	userUsecase *usecase.UserUsecase `di.inject:"userUsecase"`
}

func (u *UserController) FindAllUsers(c *fiber.Ctx) error {
	response, err := u.userUsecase.FindAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (u *UserController) FindOneUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return exception.NewHTTPError(400, "invalid user id")
	}

	response, errFind := u.userUsecase.FindOneUser(int32(userId))
	if errFind != nil {
		return errFind
	}

	return c.JSON(response)
}

func (u *UserController) FindMyUser(c *fiber.Ctx) error {
	userId := auth.GetUserIDFromClaims(c)

	response, errFind := u.userUsecase.FindOneUser(userId)
	if errFind != nil {
		return errFind
	}

	return c.JSON(response)
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
		return exception.NewValidationError(err)
	}

	response, errUpdate := u.userUsecase.UpdateUser(int32(userId), payload)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(response)
}

func (u *UserController) UpdateMyUser(c *fiber.Ctx) error {
	userId := auth.GetUserIDFromClaims(c)

	payload := new(dto.UpdateUserRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	response, errUpdate := u.userUsecase.UpdateUser(userId, payload)
	if errUpdate != nil {
		return errUpdate
	}

	return c.JSON(response)
}

func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return exception.NewHTTPError(400, "invalid user id")
	}

	response, errDelete := u.userUsecase.DeleteUser(int32(userId))
	if errDelete != nil {
		return errDelete
	}

	return c.JSON(response)
}

func (u *UserController) DeleteMyUser(c *fiber.Ctx) error {
	userId := auth.GetUserIDFromClaims(c)

	response, errDelete := u.userUsecase.DeleteUser(userId)
	if errDelete != nil {
		return errDelete
	}

	return c.JSON(response)
}
