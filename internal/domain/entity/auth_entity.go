package entity

import "time"

type Auth struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	RefreshToken string `gorm:"not null;unique"`
	ID           uint   `gorm:"primarykey"`
}
