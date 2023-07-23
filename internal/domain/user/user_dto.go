package user

import (
	"time"

	"github.com/asaskevich/govalidator"

	"codeberg.org/tfkhdyt/blog-api/pkg/validator"
)

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
}

func (r *UpdateUserRequest) Validate() (*User, error) {
	if _, err := govalidator.ValidateStruct(r); err != nil {
		return nil, validator.NewValidationError(err)
	}

	return &User{
		FullName: r.FullName,
		Username: r.Username,
		Email:    r.Email,
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
