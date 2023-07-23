package jwt

import "os"

var (
	JwtAccessTokenKey  = os.Getenv("JWT_ACCESS_TOKEN_KEY")
	JwtRefreshTokenKey = os.Getenv("JWT_REFRESH_TOKEN_KEY")
)
