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

func startFiber() {
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

	userRepo, _, userHandler := user.InjectUser(postgres.DB)
	_, _, authHandler := auth.InjectAuth(postgres.DB, userRepo)

	user.RouteUser(app, userHandler)
	auth.RouteAuth(app, authHandler)

	log.Fatalln(app.Listen(fmt.Sprintf(":%d", *port)))
}
