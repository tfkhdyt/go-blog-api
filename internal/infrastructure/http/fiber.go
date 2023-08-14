package http

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/goioc/di"
	_ "github.com/joho/godotenv/autoload"

	"codeberg.org/tfkhdyt/blog-api/internal/application/usecase"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/database"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/repository/postgres"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/security"
	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/controller"
	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/route"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
	"codeberg.org/tfkhdyt/blog-api/pkg/validator"
)

func initFiber() {
	di.RegisterBean("userRoute", reflect.TypeOf((*route.UserRoute)(nil)))
	di.RegisterBean("authRoute", reflect.TypeOf((*route.AuthRoute)(nil)))

	di.RegisterBean("userController", reflect.TypeOf((*controller.UserController)(nil)))
	di.RegisterBean("authController", reflect.TypeOf((*controller.AuthController)(nil)))

	di.RegisterBean("userUsecase", reflect.TypeOf((*usecase.UserUsecase)(nil)))
	di.RegisterBean("authUsecase", reflect.TypeOf((*usecase.AuthUsecase)(nil)))
	di.RegisterBean(
		"resetPasswordTokenUsecase",
		reflect.TypeOf((*usecase.ResetPasswordTokenUsecase)(nil)),
	)

	di.RegisterBean("userRepo", reflect.TypeOf((*postgres.UserRepositoryPostgres)(nil)))
	di.RegisterBean("authRepo", reflect.TypeOf((*postgres.AuthRepositoryPostgres)(nil)))
	di.RegisterBean(
		"resetPasswordTokenRepo",
		reflect.TypeOf((*postgres.ResetPasswordTokenRepositoryPostgres)(nil)),
	)

	di.RegisterBean("passwordHashService", reflect.TypeOf((*security.BcryptService)(nil)))
	di.RegisterBean("authTokenService", reflect.TypeOf((*security.JwtService)(nil)))
	di.RegisterBean("idService", reflect.TypeOf((*security.UUIDService)(nil)))

	di.RegisterBeanInstance("database", database.PostgresInstance)

	di.InitializeContainer()
}

func StartFiber() {
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
