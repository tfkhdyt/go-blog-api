package container

import (
	"context"
	"log"
	"reflect"

	"github.com/goioc/di"
	"github.com/resendlabs/resend-go"

	"codeberg.org/tfkhdyt/blog-api/config"
	"codeberg.org/tfkhdyt/blog-api/internal/application/usecase"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/database"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/email"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/repository/postgres"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/security"
	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/controller"
	"codeberg.org/tfkhdyt/blog-api/internal/interface/api/route"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type bean struct {
	beanType reflect.Type
	beanID   string
}

func registerBeans(beans ...bean) {
	for _, bean := range beans {
		if _, err := di.RegisterBean(bean.beanID, bean.beanType); err != nil {
			log.Fatalln("Error:", err.Error())
		}
	}
}

func InitDI() {
	registerBeans(
		bean{
			beanID:   "userRoute",
			beanType: reflect.TypeOf((*route.UserRoute)(nil)),
		},
		bean{
			beanID:   "authRoute",
			beanType: reflect.TypeOf((*route.AuthRoute)(nil)),
		},
		bean{
			beanID:   "userController",
			beanType: reflect.TypeOf((*controller.UserController)(nil)),
		},
		bean{
			beanID:   "authController",
			beanType: reflect.TypeOf((*controller.AuthController)(nil)),
		},
		bean{
			beanID:   "userUsecase",
			beanType: reflect.TypeOf((*usecase.UserUsecase)(nil)),
		},
		bean{
			beanID:   "authUsecase",
			beanType: reflect.TypeOf((*usecase.AuthUsecase)(nil)),
		},
		bean{
			beanID:   "resetPasswordUsecase",
			beanType: reflect.TypeOf((*usecase.ResetPasswordUsecase)(nil)),
		},
		bean{
			beanID:   "changeEmailUsecase",
			beanType: reflect.TypeOf((*usecase.ChangeEmailUsecase)(nil)),
		},
		bean{
			beanID:   "userRepo",
			beanType: reflect.TypeOf((*postgres.UserRepositoryPostgres)(nil)),
		},
		bean{
			beanID:   "authRepo",
			beanType: reflect.TypeOf((*postgres.AuthRepositoryPostgres)(nil)),
		},
		bean{
			beanID: "resetPasswordTokenRepo",
			beanType: reflect.TypeOf(
				(*postgres.ResetPasswordTokenRepositoryPostgres)(nil),
			),
		},
		bean{
			beanID: "changeEmailRequestRepo",
			beanType: reflect.TypeOf(
				(*postgres.ChangeEmailRequestRepositoryPostgres)(nil),
			),
		},
		bean{
			beanID:   "passwordHashService",
			beanType: reflect.TypeOf((*security.BcryptService)(nil)),
		},
		bean{
			beanID:   "authTokenService",
			beanType: reflect.TypeOf((*security.JwtService)(nil)),
		},
		bean{
			beanID:   "idService",
			beanType: reflect.TypeOf((*security.UUIDService)(nil)),
		},
		bean{
			beanID:   "emailService",
			beanType: reflect.TypeOf((*email.ResendService)(nil)),
		},
	)

	if _, err := di.RegisterBeanInstance(
		"database",
		database.PostgresInstance,
	); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if _, err := di.RegisterBeanFactory(
		"emailClient",
		di.Singleton,
		func(_ context.Context) (interface{}, error) {
			client := resend.NewClient(config.ResendApiKey)
			if client == nil {
				return nil, exception.NewHTTPError(500, "failed to initialize email client")
			}

			return client, nil
		},
	); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if err := di.InitializeContainer(); err != nil {
		log.Fatalln("Error:", err.Error())
	}
}
