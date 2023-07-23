package auth

type AuthRepository interface {
	AddToken(auth *Auth) (*Auth, error)
	VerifyToken(token string) error
	RemoveToken(token string) error
}
