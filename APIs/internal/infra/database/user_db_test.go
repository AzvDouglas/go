package database

import (
	"testing"

	"github.com/AzvDouglas/go/APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//import para user .Close()
)

func TestCreateUser(t *testing.T) {
	
	/* Outra forma de fazer o teste  
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").
		WithArgs("Douglas", "
		... 
	*/

	//Utilizando o gorm
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.User{})
	user := entity.User{
		Name:  "Douglas",
		Email: "azvdouglas@email.com",
	}

	userDB := NewUser(db)
	if err := userDB.Create(&user); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}


	assert.Nil(t, err)
	assert.Equal(t, "Douglas", user.Name)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, "Douglas", userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
