package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Douglas Azevedo", "azvdouglas@devmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Douglas Azevedo", user.Name)
	assert.Equal(t, "azvdouglas@devmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Douglas Azevedo", "azvdouglas@devmail.com", "123456")
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.Nil(t, err)
	assert.NotEqual(t, "123456", user.Password)
}