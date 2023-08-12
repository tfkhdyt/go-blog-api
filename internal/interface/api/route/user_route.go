package route

import (
	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/controller"
	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/middleware"
)

type UserRoute struct {
	userController *controller.UserController `di.inject:"userController"`
}

func (u *UserRoute) RegisterRoute(router fiber.Router) {
	router.Get("/me", middleware.JwtMiddleware, u.userController.FindMyUser)
	router.Put("/me", middleware.JwtMiddleware, u.userController.UpdateMyUser)
	router.Delete("/me", middleware.JwtMiddleware, u.userController.DeleteMyUser)

	router.Get(
		"/",
		middleware.JwtMiddleware,
		middleware.IsAdmin,
		u.userController.FindAllUsers,
	)
	router.Get(
		"/:userId",
		middleware.JwtMiddleware,
		middleware.IsAdmin,
		u.userController.FindOneUser,
	)
	router.Put(
		"/:userId",
		middleware.JwtMiddleware,
		middleware.IsAdmin,
		u.userController.UpdateUser,
	)
	router.Delete(
		"/:userId",
		middleware.JwtMiddleware,
		middleware.IsAdmin,
		u.userController.DeleteUser,
	)
}
