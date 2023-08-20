package entity

import "time"

type ChangeEmailRequest struct {
	Token     string
	NewEmail  string
	ExpiresAt time.Time
	UserID    int32
}
