package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/pkg/auth"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

var jwtAccessTokenKey = os.Getenv("JWT_ACCESS_TOKEN_KEY")

var JwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey:  jwtware.SigningKey{Key: []byte(jwtAccessTokenKey)},
	TokenLookup: "header:Authorization,cookie:accessToken",
	AuthScheme:  "Bearer",
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return exception.NewHTTPError(401, err.Error())
	},
})

var IsAdmin = func(c *fiber.Ctx) error {
	role := auth.GetRoleFromClaims(c)

	if role != "admin" {
		return exception.NewHTTPError(403, "you're not allowed to access this endpoint")
	}

	return c.Next()
}
