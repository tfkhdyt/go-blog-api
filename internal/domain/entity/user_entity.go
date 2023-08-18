package entity

import "time"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	FullName  string
	Username  string
	Email     string
	Password  string
	Role      Role
	ID        int32
}
