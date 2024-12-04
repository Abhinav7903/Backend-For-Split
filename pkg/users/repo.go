package users

import (
	"time"

	"github.com/Abhinav7903/split/factory"
)

type Repository interface {
	AddUser(user factory.User) error
	VerifyEmail(email string) error
	LoginUser(user factory.User) (string, bool, string, time.Time, error)
}
