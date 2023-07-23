package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/auth"
)

type authHandler struct {
	authService auth.AuthService
}

func NewAuthHandler(authService auth.AuthService) *authHandler {
	return &authHandler{authService}
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
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

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
