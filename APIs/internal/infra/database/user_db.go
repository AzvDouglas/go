package database

import (
	"gorm.io/gorm"
	"github.com/AzvDouglas/go/APIs/internal/entity"
)

type User struct {
	DB *gorm.DB
}

//func (p *User) Create(User *entity.User) error {
func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(User *entity.User) error {
	return u.DB.Create(User).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
