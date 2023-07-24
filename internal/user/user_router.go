package user

import (
	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/auth"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

func RouteUser(app *fiber.App, userHandler user.UserHandler) {
	app.Get("/users/me", auth.JwtMiddleware, userHandler.FindMyUser)
	app.Put("/users/me", auth.JwtMiddleware, userHandler.UpdateMyUser)
	app.Delete("/users/me", auth.JwtMiddleware, userHandler.DeleteMyUser)

	app.Get("/users", auth.JwtMiddleware, auth.IsAdmin, userHandler.FindAllUsers)
	app.Get("/users/:userId", auth.JwtMiddleware, auth.IsAdmin, userHandler.FindOneUser)
	app.Put("/users/:userId", auth.JwtMiddleware, auth.IsAdmin, userHandler.UpdateUser)
	app.Delete("/users/:userId", auth.JwtMiddleware, auth.IsAdmin, userHandler.DeleteUser)
}
