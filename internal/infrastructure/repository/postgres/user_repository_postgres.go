package postgres

import (
	"fmt"

	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type UserRepositoryPostgres struct {
	db *gorm.DB `di.inject:"database"`
}

func (u *UserRepositoryPostgres) Register(newUser *entity.User) (*entity.User, error) {
	if err := u.db.Create(newUser).Error; err != nil {
		return nil, exception.NewHTTPError(500, "failed to register new user")
	}

	return newUser, nil
}

func (u *UserRepositoryPostgres) FindAllUsers() (*[]entity.User, error) {
	users := new([]entity.User)
	if err := u.db.Order("id").Find(users).Error; err != nil {
		return nil, exception.NewHTTPError(500, "failed to find all users")
	}

	return users, nil
}

func (u *UserRepositoryPostgres) FindOneUser(userId uint) (*entity.User, error) {
	user := new(entity.User)
	if err := u.db.First(user, userId).Error; err != nil {
		return nil, exception.NewHTTPError(404, fmt.Sprintf("user with id %d is not found", userId))
	}

	return user, nil
}

func (u *UserRepositoryPostgres) FindOneUserByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	if err := u.db.First(user, "email = ?", email).Error; err != nil {
		return nil, exception.NewHTTPError(
			404,
			fmt.Sprintf("user with email %s is not found", email),
		)
	}

	return user, nil
}

func (u *UserRepositoryPostgres) UpdateUser(
	oldUser *entity.User,
	newUser *entity.User,
) (*entity.User, error) {
	if err := u.db.Model(oldUser).Updates(newUser).Error; err != nil {
		return nil, exception.NewHTTPError(
			500,
			fmt.Sprintf("failed to update user with id %d", oldUser.ID),
		)
	}

	return oldUser, nil
}

func (u *UserRepositoryPostgres) DeleteUser(userId uint) error {
	if err := u.db.Delete(&entity.User{}, userId).Error; err != nil {
		return exception.NewHTTPError(
			500,
			fmt.Sprintf("failed to delete user with id %d", userId),
		)
	}

	return nil
}
