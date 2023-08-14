package service

type AuthTokenService interface {
	CreateAccessToken(id uint, role string) (string, error)
	CreateRefreshToken(id uint, role string) (string, error)
}
