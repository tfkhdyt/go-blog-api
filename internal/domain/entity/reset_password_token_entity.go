package entity

import "time"

type ResetPasswordToken struct {
	ExpiresAt time.Time
	Token     string `gorm:"not null;unique"`
	UserID    uint
	ID        uint `gorm:"primarykey"`
}
