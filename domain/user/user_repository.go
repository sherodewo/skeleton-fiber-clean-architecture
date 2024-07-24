package user

type UserRepository interface {
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
}
