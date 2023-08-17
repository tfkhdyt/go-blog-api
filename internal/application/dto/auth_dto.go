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

type RegisterResponseData struct {
	CreatedAt time.Time `json:"created_at"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	ID        uint      `json:"id"`
}

type RegisterResponse struct {
	Message string               `json:"message"`
	Data    RegisterResponseData `json:"data"`
}

// Login
type LoginRequest struct {
	Email    string `json:"email"    valid:"required~email is required,email~invalid email"`
	Password string `json:"password" valid:"required~password is required,stringlength(8|128)~password length should be between 8 - 128 chars"`
}

type LoginResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
}

// Refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" valid:"required~refresh token is required"`
}

type RefreshResponseData struct {
	AccessToken string `json:"access_token"`
}

type RefreshResponse struct {
	Message string              `json:"message"`
	Data    RefreshResponseData `json:"data"`
}

// Logout
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" valid:"required~refresh token is required"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

// Change password
type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password"     valid:"required~old password is required,stringlength(8|128)~old password length should be between 8 - 128 chars"`
	NewPassword     string `json:"new_password"     valid:"required~new password is required,stringlength(8|128)~new password length should be between 8 - 128 chars"`
	ConfirmPassword string `json:"confirm_password" valid:"required~confim password is required,stringlength(8|128)~confirm password length should be between 8 - 128 chars"`
}

type ChangePasswordResponse LogoutResponse
