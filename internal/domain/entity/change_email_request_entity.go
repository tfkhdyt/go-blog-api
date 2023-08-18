package entity

import "time"

type ChangeEmailRequest struct {
	ExpiresAt time.Time
	Token     string
	NewEmail  string
	UserID    int32
}
