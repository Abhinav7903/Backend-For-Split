package users

import (
	"github.com/Abhinav7903/split/factory"
)

type Repository interface {
	AddUser(user factory.User) error
	VerifyEmail(email string) error
	GetUser(email string) (factory.User, error)
	UpdateUserDetails(user factory.User) error
	DeleteUser(email string) error
	GetAllUsers() ([]factory.User, error)
	EmailExists(email string) (bool, error)
}
