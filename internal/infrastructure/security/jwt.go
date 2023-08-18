package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"codeberg.org/tfkhdyt/blog-api/config"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/service"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type JwtService struct{}

func (j *JwtService) CreateAccessToken(id int32, role entity.Role) (string, error) {
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

func (j *JwtService) CreateRefreshToken(
	id int32,
	role entity.Role,
) (string, error) {
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

func (j *JwtService) ParseRefreshToken(tokenString string) (*service.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.NewHTTPError(
				400,
				fmt.Sprintf(
					"unexpected signing method: %v",
					token.Header["alg"],
				),
			)
		}

		return []byte(config.JwtRefreshTokenKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, exception.NewHTTPError(400, "invalid token")
	}

	userId, okUserId := claims["userId"].(float64)
	if !okUserId {
		return nil, exception.
			NewHTTPError(400, "failed to parse user id from claims")
	}

	role, okRole := claims["role"].(string)
	if !okRole {
		return nil, exception.NewHTTPError(400, "failed to parse role from claims")
	}

	return &service.Claims{
		UserID: int32(userId),
		Role:   entity.Role(role),
	}, nil
}
