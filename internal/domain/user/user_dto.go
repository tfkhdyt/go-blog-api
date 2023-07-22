package user

import (
	"time"

	"github.com/asaskevich/govalidator"

	"codeberg.org/tfkhdyt/blog-api/pkg/validator"
)

// Register
type RegisterRequest struct {
	FullName string `json:"full_name" valid:"required~full name is required,stringlength(3|50)~full name length should be between 3 - 50 chars"`
	Username string `json:"username"  valid:"required~username is required,stringlength(3|16)~username length should be between 3 - 16 chars"`
	Email    string `json:"email"     valid:"required~email is required,email~invalid email"`
	Password string `json:"password"  valid:"required~password is required,stringlength(8|128)~password length should be between 8 - 128 chars"`
}

func (r *RegisterRequest) Validate() (*User, error) {
	if _, err := govalidator.ValidateStruct(r); err != nil {
		return nil, validator.NewValidationError(err)
	}

	return &User{
		FullName: r.FullName,
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Role:     "user",
	}, nil
}

type RegisterResponse struct {
	CreatedAt time.Time `json:"created_at"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	ID        uint      `json:"id"`
}

// Find all
type FindAllUsersResponse []FindOneUserResponse

// Find one
type FindOneUserResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	ID        uint      `json:"id"`
}

// Update
type UpdateUserRequest struct {
	FullName string `json:"full_name" valid:"stringlength(3|50)~full name length should be between 3 - 50 chars"`
	Username string `json:"username"  valid:"stringlength(3|16)~username length should be between 3 - 16 chars"`
	Email    string `json:"email"     valid:"email~invalid email"`
	Password string `json:"password"  valid:"stringlength(8|128)~password length should be between 8 - 128 chars"`
}

func (r *UpdateUserRequest) Validate() (*User, error) {
	if _, err := govalidator.ValidateStruct(r); err != nil {
		return nil, validator.NewValidationError(err)
	}

	return &User{
		FullName: r.FullName,
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}, nil
}

type UpdateUserResponse struct {
	UpdatedAt time.Time `json:"updated_at"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	ID        uint      `json:"id"`
}

// Delete
type DeleteUserResponse struct {
	Message string `json:"message"`
}
