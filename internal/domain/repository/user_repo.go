package repository

import "codeberg.org/tfkhdyt/blog-api/internal/domain/entity"

type UserRepository interface {
	Register(newUser *entity.User) (*entity.User, error)
	FindAllUsers() (*[]entity.User, error)
	FindOneUser(userId uint) (*entity.User, error)
	FindOneUserByEmail(email string) (*entity.User, error)
	UpdateUser(oldUser *entity.User, newUser *entity.User) (*entity.User, error)
	DeleteUser(userId uint) error
}
