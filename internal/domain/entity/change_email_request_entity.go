package entity

import "time"

type ChangeEmailRequest struct {
	ExpiresAt time.Time
	Token     string `gorm:"not null;unique"`
	NewEmail  string `gorm:"not null"`
	UserID    uint
	ID        uint `gorm:"primarykey"`
}
