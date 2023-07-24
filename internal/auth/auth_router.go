package auth

import (
	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/auth"
)

func RouteAuth(app *fiber.App, authHandler auth.AuthHandler) {
	app.Post("/auth/register", authHandler.Register)
	app.Post("/auth/login", authHandler.Login)

	app.Put("/auth/refresh", JwtMiddleware, authHandler.Refresh)
	app.Delete("/auth/logout", JwtMiddleware, authHandler.Logout)
}
