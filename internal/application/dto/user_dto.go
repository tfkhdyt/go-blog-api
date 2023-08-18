package dto

import (
	"time"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
)

// Find all
type FindAllUsersResponseData []FindOneUserResponseData

type FindAllUsersResponse struct {
	Data FindAllUsersResponseData `json:"data"`
}

// Find one
type FindOneUserResponseData struct {
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	FullName  string      `json:"full_name"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Role      entity.Role `json:"role"`
	ID        int32       `json:"id"`
}

type FindOneUserResponse struct {
	Data FindOneUserResponseData `json:"data"`
}

// Update
type UpdateUserRequest struct {
	FullName string `json:"full_name" valid:"stringlength(3|50)~full name length should be between 3 - 50 chars"`
	Username string `json:"username"  valid:"stringlength(3|16)~username length should be between 3 - 16 chars"`
}

type UpdateUserResponseData struct {
	UpdatedAt time.Time   `json:"updated_at"`
	FullName  string      `json:"full_name"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Role      entity.Role `json:"role"`
	ID        int32       `json:"id"`
}

type UpdateUserResponse struct {
	Message string                 `json:"message"`
	Data    UpdateUserResponseData `json:"data"`
}

// Delete
type DeleteUserResponse struct {
	Message string `json:"message"`
}
