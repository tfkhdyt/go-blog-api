package route

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/goioc/di"

	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/controller"
	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/middleware"
)

type AuthRoute struct {
	authController *controller.AuthController `di.inject:"authController"`
}

func (a *AuthRoute) RegisterRoute(router fiber.Router) {
	router.Post("/register", a.authController.Register)
	router.Post("/login", a.authController.Login)

	router.Put("/refresh", middleware.JwtMiddleware, a.authController.Refresh)
	router.Delete("/logout", middleware.JwtMiddleware, a.authController.Logout)
}
