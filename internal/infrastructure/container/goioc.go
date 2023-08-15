package container

import (
	"reflect"

	"github.com/goioc/di"

	"codeberg.org/tfkhdyt/blog-api/internal/application/usecase"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/database"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/repository/postgres"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/security"
	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/controller"
	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/route"
)

func InitDI() {
	// routes
	di.RegisterBean("userRoute", reflect.TypeOf((*route.UserRoute)(nil)))
	di.RegisterBean("authRoute", reflect.TypeOf((*route.AuthRoute)(nil)))

	// controllers
	di.RegisterBean("userController", reflect.TypeOf((*controller.UserController)(nil)))
	di.RegisterBean("authController", reflect.TypeOf((*controller.AuthController)(nil)))

	// usecases
	di.RegisterBean("userUsecase", reflect.TypeOf((*usecase.UserUsecase)(nil)))
	di.RegisterBean("authUsecase", reflect.TypeOf((*usecase.AuthUsecase)(nil)))
	di.RegisterBean(
		"resetPasswordTokenUsecase",
		reflect.TypeOf((*usecase.ResetPasswordTokenUsecase)(nil)),
	)

	// repositories
	di.RegisterBean("userRepo", reflect.TypeOf((*postgres.UserRepositoryPostgres)(nil)))
	di.RegisterBean("authRepo", reflect.TypeOf((*postgres.AuthRepositoryPostgres)(nil)))
	di.RegisterBean(
		"resetPasswordTokenRepo",
		reflect.TypeOf((*postgres.ResetPasswordTokenRepositoryPostgres)(nil)),
	)

	// services
	di.RegisterBean("passwordHashService", reflect.TypeOf((*security.BcryptService)(nil)))
	di.RegisterBean("authTokenService", reflect.TypeOf((*security.JwtService)(nil)))
	di.RegisterBean("idService", reflect.TypeOf((*security.UUIDService)(nil)))

	// databases
	di.RegisterBeanInstance("database", database.PostgresInstance)

	di.InitializeContainer()
}
