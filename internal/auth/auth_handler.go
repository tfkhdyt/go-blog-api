package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/auth"
	authHelper "codeberg.org/tfkhdyt/blog-api/pkg/auth"
)

type authHandler struct {
	authService auth.AuthService
}

func NewAuthHandler(authService auth.AuthService) *authHandler {
	return &authHandler{authService}
}

func (a *authHandler) Register(c *fiber.Ctx) error {
	payload := new(auth.RegisterRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "failed to parse body")
	}

	registeredUser, err := a.authService.Register(payload)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": registeredUser,
	})
}

func (a *authHandler) Login(c *fiber.Ctx) error {
	auth_ := new(auth.LoginRequest)
	if err := c.BodyParser(auth_); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "failed to parse body")
	}

	response, err := a.authService.Login(auth_)
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

func (a *authHandler) Refresh(c *fiber.Ctx) error {
	userId := authHelper.GetUserIDFromClaims(c)

	ath := new(auth.RefreshRequest)
	if err := c.BodyParser(ath); err != nil {
		rfrsh := c.Cookies("refreshToken")
		if rfrsh == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid refresh token")
		}

		ath.RefreshToken = rfrsh
	}

	response, err := a.authService.Refresh(uint(userId), ath)
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

func (a *authHandler) Logout(c *fiber.Ctx) error {
	ath := new(auth.LogoutRequest)
	if err := c.BodyParser(ath); err != nil {
		rfrsh := c.Cookies("refreshToken")
		if rfrsh == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid refresh token")
		}

		ath.RefreshToken = rfrsh
	}

	response, err := a.authService.Logout(ath.RefreshToken)
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
