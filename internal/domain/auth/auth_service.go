package auth

type AuthService interface {
	Register(payload *RegisterRequest) (*RegisterResponse, error)
	Login(payload *LoginRequest) (*LoginResponse, error)
	Refresh(userId uint, payload *RefreshRequest) (*RefreshResponse, error)
	Logout(refreshToken string) (*LogoutResponse, error)
}
