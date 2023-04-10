package entity

import (
	// "github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
	"time"

	"github.com/AzvDouglas/go/APIs/pkg/entity"
	"github.com/pkg/errors"
)

var (
	ErrorIDIsRequired = errors.New("ID is required")
	ErrorNameIsRequired = errors.New("Name is required")
	ErrorDescriptionIsRequired = errors.New("Description is required")
	ErrorPriceIsRequired = errors.New("Price is required")
	ErrorInvalidPrice = errors.New("Invalid price")
)

type Product struct {
	ID          entity.ID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorIDIsRequired
	}

	if p.Name == "" {
		return ErrorNameIsRequired
	}
	
	if p.Price == 0 {
		return ErrorPriceIsRequired
	}
	if p.Price < 0 {
		return ErrorInvalidPrice
	}
	return nil
}

func NewProduct(name, description string, price float64) (*Product, error) {
	product := &Product{
		ID:          	entity.NewID(),
		Name:        	name,
		Price:			price,
		Description: 	description,
		CreatedAt:  	time.Now(),
	}
	err := product.Validate() 
	if err != nil {
		return nil, err
	}
	return product, nil
}