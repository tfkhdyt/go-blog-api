package user

import (
	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

type UserInjector struct {
	UserRepo    user.UserRepository
	UserHandler user.UserHandler
}

func InjectUser(db *gorm.DB) *UserInjector {
	userRepo := NewUserRepoPostgres(db)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	return &UserInjector{userRepo, userHandler}
}
