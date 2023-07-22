package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null;size:50"`
	Username string `gorm:"not null;unique;size:16"`
	Email    string `gorm:"not null;unique;size:25"`
	Password string `gorm:"not null;size:128"`
	Role     string `gorm:"not null;default:user;size:10"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return fmt.Errorf("failed to hash Password")
	}

	u.Password = string(hashedPassword)

	return nil
}
