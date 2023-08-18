package service

import "codeberg.org/tfkhdyt/blog-api/internal/domain/entity"

type Claims struct {
	Role   entity.Role
	UserID int32
}

type AuthTokenService interface {
	CreateAccessToken(id int32, role entity.Role) (string, error)
	CreateRefreshToken(id int32, role entity.Role) (string, error)
	ParseRefreshToken(tokenString string) (*Claims, error)
}
