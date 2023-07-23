package auth

import (
	"github.com/asaskevich/govalidator"

	"codeberg.org/tfkhdyt/blog-api/pkg/validator"
)

// Login
type LoginRequest struct {
	Email    string `json:"email"    valid:"required~email is required,email~invalid email"`
	Password string `json:"password" valid:"required~password is required,stringlength(8|128)~password length should be between 8 - 128 chars"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (l *LoginRequest) Validate() error {
	if _, err := govalidator.ValidateStruct(l); err != nil {
		return validator.NewValidationError(err)
	}

	return nil
}

// Refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" valid:"required~refresh token is required"`
}

func (r *RefreshRequest) Validate() error {
	if _, err := govalidator.ValidateStruct(r); err != nil {
		return validator.NewValidationError(err)
	}

	return nil
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

// Logout
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" valid:"required~refresh token is required"`
}

func (l *LogoutRequest) Validate() error {
	if _, err := govalidator.ValidateStruct(l); err != nil {
		return validator.NewValidationError(err)
	}

	return nil
}

type LogoutResponse struct {
	Message string `json:"message"`
}
