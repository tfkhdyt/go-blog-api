package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"codeberg.org/tfkhdyt/blog-api/config"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type JwtService struct{}

func (j *JwtService) CreateAccessToken(id uint, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    "blog-api",
		"userId": id,
		"role":   role,
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	})
	if token == nil {
		return "", exception.
			NewHTTPError(500, "failed to to create new access token")
	}

	signedString, err := token.SignedString([]byte(config.JwtAccessTokenKey))
	if err != nil {
		return "", exception.NewHTTPError(500, "failed to sign access token")
	}

	return signedString, nil
}

func (j *JwtService) CreateRefreshToken(id uint, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    "blog-api",
		"userId": id,
		"role":   role,
		"exp":    time.Now().Add(720 * time.Hour).Unix(),
	})
	if token == nil {
		return "", exception.
			NewHTTPError(500, "failed to to create new refresh token")
	}

	signedString, err := token.SignedString([]byte(config.JwtRefreshTokenKey))
	if err != nil {
		return "", exception.NewHTTPError(500, "failed to sign refresh token")
	}

	return signedString, nil
}
