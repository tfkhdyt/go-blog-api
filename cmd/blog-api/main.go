package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"

	"codeberg.org/tfkhdyt/blog-api/internal/auth"
	"codeberg.org/tfkhdyt/blog-api/internal/database/postgres"
	"codeberg.org/tfkhdyt/blog-api/internal/user"
	"codeberg.org/tfkhdyt/blog-api/pkg/validator"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			var valErr *validator.Error
			if errors.As(err, &valErr) {
				errs := strings.Split(err.Error(), ";")
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"errors": errs,
				})
			}

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	app.Use(recover.New())
	port := flag.Uint("port", 8080, "server port")
	flag.Parse()

	userRepo := user.NewUserRepoPostgres(postgres.DB)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	authRepo := auth.NewAuthRepoPostgres(postgres.DB)
	authService := auth.NewAuthService(authRepo, userRepo)
	authHandler := auth.NewAuthHandler(authService)

	app.Get("/users", auth.JwtMiddleware, auth.IsAdmin, userHandler.FindAllUsers)
	app.Get("/users/me", auth.JwtMiddleware, userHandler.FindMyUser)
	app.Get("/users/:userId", auth.JwtMiddleware, auth.IsAdmin, userHandler.FindOneUser)
	app.Put("/users/me", auth.JwtMiddleware, userHandler.UpdateMyUser)
	app.Put("/users/:userId", auth.JwtMiddleware, auth.IsAdmin, userHandler.UpdateUser)
	app.Delete("/users/me", auth.JwtMiddleware, userHandler.DeleteMyUser)
	app.Delete("/users/:userId", auth.JwtMiddleware, auth.IsAdmin, userHandler.DeleteUser)

	app.Post("/auth/register", authHandler.Register)
	app.Post("/auth/login", authHandler.Login)
	app.Put("/auth/refresh", auth.JwtMiddleware, authHandler.Refresh)
	app.Delete("/auth/logout", auth.JwtMiddleware, authHandler.Logout)

	app.Get("/restricted", auth.JwtMiddleware, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusTeapot).JSON(fiber.Map{
			"hello": "world!",
		})
	})

	log.Fatalln(
		app.Listen(
			fmt.Sprintf(":%d", *port),
		),
	)
}
