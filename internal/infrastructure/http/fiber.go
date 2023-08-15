package http

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/goioc/di"

	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/route"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
	"codeberg.org/tfkhdyt/blog-api/pkg/validator"
)

func StartFiberServer() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *exception.HttpError
			if errors.As(err, &e) {
				code = int(e.StatusCode())
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
	app.Use(pprof.New())

	port := flag.Uint("port", 8080, "server port")
	flag.Parse()

	di.GetInstance("userRoute").(*route.UserRoute).
		RegisterRoute(app.Group("/users"))
	di.GetInstance("authRoute").(*route.AuthRoute).
		RegisterRoute(app.Group("/auth"))

	log.Fatalln(app.Listen(fmt.Sprintf(":%d", *port)))
}
