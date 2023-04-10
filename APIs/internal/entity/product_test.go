package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", "lorem description", 10.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, "lorem description", product.Description)
	assert.Equal(t, 10.00, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", "lorem description", 10.00)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrorNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", "lorem description", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrorPriceIsRequired, err)
}

func TestProduct_Validate(t *testing.T) {
	product, err := NewProduct("Product 1", "lorem description", 10.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, "lorem description", product.Description)
	assert.Equal(t, 10.00, product.Price)
}
