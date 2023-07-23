package auth

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	RefreshToken string `gorm:"not null;unique"`
}
