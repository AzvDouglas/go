package database

import (
	"github.com/AzvDouglas/go/APIs/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
