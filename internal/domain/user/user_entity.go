package user

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	jwtConfig "codeberg.org/tfkhdyt/blog-api/internal/config/jwt"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null;size:50"`
	Username string `gorm:"not null;unique;size:16"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null;default:user;size:10"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return fmt.Errorf("failed to hash Password")
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}

func (u *User) CreateAccessToken() (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    "blog-api",
		"userId": u.ID,
		"role":   u.Role,
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	})
	if token == nil {
		return nil, fmt.Errorf("failed to to create new access token")
	}

	signedString, err := token.SignedString([]byte(jwtConfig.JwtAccessTokenKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token")
	}

	return &signedString, nil
}

func (u *User) CreateRefreshToken() (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    "blog-api",
		"userId": u.ID,
		"role":   u.Role,
		"exp":    time.Now().Add(720 * time.Hour).Unix(),
	})
	if token == nil {
		return nil, fmt.Errorf("failed to to create new refresh token")
	}

	signedString, err := token.SignedString([]byte(jwtConfig.JwtRefreshTokenKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token")
	}

	return &signedString, nil
}
