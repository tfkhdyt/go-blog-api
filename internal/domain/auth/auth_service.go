package auth

type AuthService interface {
	Login(payload *LoginRequest) (*LoginResponse, error)
	Refresh(userId uint, payload *RefreshRequest) (*RefreshResponse, error)
	Logout(refreshToken string) (*LogoutResponse, error)
}
