package service

type PasswordHashService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) error
}
