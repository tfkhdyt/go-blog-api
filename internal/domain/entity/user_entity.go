package entity

import "time"

type User struct {
	CreatedAt           time.Time
	UpdatedAt           time.Time
	FullName            string               `gorm:"not null;size:50"`
	Username            string               `gorm:"not null;unique;size:16"`
	Email               string               `gorm:"not null;unique"`
	Password            string               `gorm:"not null"`
	Role                string               `gorm:"not null;default:user;size:10"`
	ResetPasswordTokens []ResetPasswordToken `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChangeEmailRequests []ChangeEmailRequest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ID                  uint                 `gorm:"primarykey"`
}
