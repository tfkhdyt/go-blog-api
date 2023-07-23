package auth

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	jwtConfig "codeberg.org/tfkhdyt/blog-api/internal/config/jwt"
	"codeberg.org/tfkhdyt/blog-api/pkg/auth"
)

var JwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey:  jwtware.SigningKey{Key: []byte(jwtConfig.JwtAccessTokenKey)},
	TokenLookup: "header:Authorization,cookie:accessToken",
	AuthScheme:  "Bearer",
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	},
})

var IsAdmin = func(c *fiber.Ctx) error {
	role := auth.GetRoleFromClaims(c)

	if role != "admin" {
		return fiber.NewError(fiber.StatusForbidden, "you're not allowed to access this endpoint")
	}

	return c.Next()
}
