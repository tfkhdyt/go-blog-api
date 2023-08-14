package security

import "github.com/google/uuid"

type UUIDService struct{}

func (u *UUIDService) GenerateID() string {
	return uuid.NewString()
}
