package user

import (
	"fmt"

	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

type userRepoPostgres struct {
	db *gorm.DB
}

func NewUserRepoPostgres(db *gorm.DB) *userRepoPostgres {
	return &userRepoPostgres{db}
}

func (u *userRepoPostgres) Register(newUser *user.User) (*user.User, error) {
	if err := u.db.Create(newUser).Error; err != nil {
		return nil, fmt.Errorf("failed to register new user")
	}

	return newUser, nil
}

func (u *userRepoPostgres) FindAllUsers() (*[]user.User, error) {
	users := new([]user.User)
	if err := u.db.Order("id").Find(users).Error; err != nil {
		return nil, fmt.Errorf("failed to find all users")
	}

	return users, nil
}

func (u *userRepoPostgres) FindOneUser(userId uint) (*user.User, error) {
	user := new(user.User)
	if err := u.db.First(user, userId).Error; err != nil {
		return nil, fmt.Errorf("user with id %d is not found", userId)
	}

	return user, nil
}

func (u *userRepoPostgres) FindOneUserByEmail(email string) (*user.User, error) {
	user := new(user.User)
	if err := u.db.First(user, "email = ?", email).Error; err != nil {
		return nil, fmt.Errorf("user with email %s is not found", email)
	}

	return user, nil
}

func (u *userRepoPostgres) UpdateUser(oldUser *user.User, newUser *user.User) (*user.User, error) {
	if err := u.db.Model(oldUser).Updates(newUser).Error; err != nil {
		return nil, fmt.Errorf("failed to update user with id %d", oldUser.ID)
	}

	return oldUser, nil
}

func (u *userRepoPostgres) DeleteUser(userId uint) error {
	if err := u.db.Delete(&user.User{}, userId).Error; err != nil {
		return fmt.Errorf("failed to delete user with id %d", userId)
	}

	return nil
}
