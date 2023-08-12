package dto

import (
	"time"
)

// Register
type RegisterRequest struct {
	FullName string `json:"full_name" valid:"required~full name is required,stringlength(3|50)~full name length should be between 3 - 50 chars"`
	Username string `json:"username"  valid:"required~username is required,stringlength(3|16)~username length should be between 3 - 16 chars"`
	Email    string `json:"email"     valid:"required~email is required,email~invalid email"`
	Password string `json:"password"  valid:"required~password is required,stringlength(8|128)~password length should be between 8 - 128 chars"`
}

// func (r *RegisterRequest) Validate() (*user.User, error) {
// 	if _, err := govalidator.ValidateStruct(r); err != nil {
// 		return nil, validator.NewValidationError(err)
// 	}
//
// 	return &user.User{
// 		FullName: r.FullName,
// 		Username: r.Username,
// 		Email:    r.Email,
// 		Password: r.Password,
// 		Role:     "user",
// 	}, nil
// }

type RegisterResponse struct {
	CreatedAt time.Time `json:"created_at"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	ID        uint      `json:"id"`
}

// Login
type LoginRequest struct {
	Email    string `json:"email"    valid:"required~email is required,email~invalid email"`
	Password string `json:"password" valid:"required~password is required,stringlength(8|128)~password length should be between 8 - 128 chars"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// func (l *LoginRequest) Validate() error {
// 	if _, err := govalidator.ValidateStruct(l); err != nil {
// 		return validator.NewValidationError(err)
// 	}
//
// 	return nil
// }

// Refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" valid:"required~refresh token is required"`
}

// func (r *RefreshRequest) Validate() error {
// 	if _, err := govalidator.ValidateStruct(r); err != nil {
// 		return validator.NewValidationError(err)
// 	}
//
// 	return nil
// }

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

// Logout
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" valid:"required~refresh token is required"`
}

// func (l *LogoutRequest) Validate() error {
// 	if _, err := govalidator.ValidateStruct(l); err != nil {
// 		return validator.NewValidationError(err)
// 	}
//
// 	return nil
// }

type LogoutResponse struct {
	Message string `json:"message"`
}
