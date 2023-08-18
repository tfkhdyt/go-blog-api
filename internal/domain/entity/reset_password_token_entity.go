package entity

import "time"

type ResetPasswordToken struct {
	Token     string
	ExpiresAt time.Time
	UserID    int32
}
