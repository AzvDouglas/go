package entity

import (
	"github.com/AzvDouglas/go/APIs/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID   `json:"id"`
	Name	 string 	 `json:"username"`
	Email    string 	 `json:"email"`
	Password string 	 `json:"_"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID: entity.NewID(),
		Name: name,
		Email: email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	//return bycrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}