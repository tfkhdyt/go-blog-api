package user

import (
	"gorm.io/gorm"
)

func InjectUser(db *gorm.DB) (*userRepoPostgres, *userService, *userHandler) {
	userRepo := NewUserRepoPostgres(db)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	return userRepo, userService, userHandler
}
