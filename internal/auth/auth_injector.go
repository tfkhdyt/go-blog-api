package auth

import (
	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/auth"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

type AuthInjector struct {
	AuthRepo    auth.AuthRepository
	AuthHandler auth.AuthHandler
}

func InjectAuth(db *gorm.DB, userRepo user.UserRepository) *AuthInjector {
	authRepo := NewAuthRepoPostgres(db)
	authService := NewAuthService(authRepo, userRepo)
	authHandler := NewAuthHandler(authService)

	return &AuthInjector{authRepo, authHandler}
}
