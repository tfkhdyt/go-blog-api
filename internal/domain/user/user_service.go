package user

type UserService interface {
	Register(payload *RegisterRequest) (*RegisterResponse, error)
	FindAllUsers() (*FindAllUsersResponse, error)
	FindOneUser(userId uint) (*FindOneUserResponse, error)
	UpdateUser(userId uint, payload *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(userId uint) (*DeleteUserResponse, error)
}
