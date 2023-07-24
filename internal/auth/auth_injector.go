package auth

import (
	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

func InjectAuth(
	db *gorm.DB,
	userRepo user.UserRepository,
) (*authRepoPostgres, *authService, *authHandler) {
	authRepo := NewAuthRepoPostgres(db)
	authService := NewAuthService(authRepo, userRepo)
	authHandler := NewAuthHandler(authService)

	return authRepo, authService, authHandler
}
