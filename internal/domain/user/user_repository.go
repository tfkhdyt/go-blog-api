package user

type UserRepository interface {
	Register(newUser *User) (*User, error)
	FindAllUsers() (*[]User, error)
	FindOneUser(userId uint) (*User, error)
	UpdateUser(oldUser *User, newUser *User) (*User, error)
	DeleteUser(userId uint) error
}
