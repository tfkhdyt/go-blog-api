package controller

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/application/usecase"
	"codeberg.org/tfkhdyt/blog-api/pkg/auth"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type AuthController struct {
	authUsecase               *usecase.AuthUsecase               `di.inject:"authUsecase"`
	resetPasswordTokenUsecase *usecase.ResetPasswordTokenUsecase `di.inject:"resetPasswordTokenUsecase"`
}

func (a *AuthController) Register(c *fiber.Ctx) error {
	payload := new(dto.RegisterRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	registeredUser, err := a.authUsecase.Register(payload)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": registeredUser,
	})
}

func (a *AuthController) Login(c *fiber.Ctx) error {
	payload := new(dto.LoginRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	response, err := a.authUsecase.Login(payload)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    response.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    response.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(720 * time.Hour),
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": response,
	})
}

func (a *AuthController) Refresh(c *fiber.Ctx) error {
	userId := auth.GetUserIDFromClaims(c)

	payload := new(dto.RefreshRequest)
	if err := c.BodyParser(payload); err != nil {
		rfrsh := c.Cookies("refreshToken")
		if rfrsh == "" {
			return exception.NewHTTPError(401, "invalid refresh token")
		}

		payload.RefreshToken = rfrsh
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	response, err := a.authUsecase.Refresh(uint(userId), payload)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    response.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"data": response,
	})
}

func (a *AuthController) Logout(c *fiber.Ctx) error {
	payload := new(dto.LogoutRequest)
	if err := c.BodyParser(payload); err != nil {
		rfrsh := c.Cookies("refreshToken")
		if rfrsh == "" {
			return exception.NewHTTPError(401, "invalid refresh token")
		}

		payload.RefreshToken = rfrsh
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	response, err := a.authUsecase.Logout(payload.RefreshToken)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    "deleted",
		Path:     "/",
		Expires:  time.Date(2002, time.April, 1, 23, 0, 0, 0, time.UTC),
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    "deleted",
		Path:     "/",
		Expires:  time.Date(2002, time.April, 1, 23, 0, 0, 0, time.UTC),
		HTTPOnly: true,
	})

	return c.JSON(response)
}

func (a *AuthController) ChangePassword(c *fiber.Ctx) error {
	userId := auth.GetUserIDFromClaims(c)

	payload := new(dto.ChangePasswordRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	response, errChangePassword := a.authUsecase.ChangePassword(userId, payload)
	if errChangePassword != nil {
		return errChangePassword
	}

	return c.JSON(response)
}

func (a *AuthController) GetResetPasswordToken(c *fiber.Ctx) error {
	payload := new(dto.GetResetPasswordTokenRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	response, errToken := a.resetPasswordTokenUsecase.
		GetResetPasswordToken(payload)
	if errToken != nil {
		return errToken
	}

	return c.Status(202).JSON(response)
}

func (a *AuthController) ResetPassword(c *fiber.Ctx) error {
	token := c.Params("token")

	payload := new(dto.ResetPasswordRequest)
	if err := c.BodyParser(payload); err != nil {
		return exception.NewHTTPError(422, "failed to parse body")
	}

	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	response, err := a.resetPasswordTokenUsecase.ResetPassword(token, payload)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
