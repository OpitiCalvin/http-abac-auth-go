package entity_test

import (
	"testing"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/stretchr/testify/assert"
)

// TestNewProduct tests creation of a new product with name and base url provided
func TestNewProduct(t *testing.T) {
	p, err := entity.NewProduct("Test Product", "http://test-product.com")

	assert.Nil(t, err)
	assert.Equal(t, p.ProductName, "Test Product")
	assert.Equal(t, p.BaseURL, "http://test-product.com")
}

// TestProductValidateNoName test validation method for a product entity
func TestProductValidateNoName(t *testing.T) {
	p, err := entity.NewProduct("", "http://test-product.com")

	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, err, entity.ErrInvalidEntity)
}

// TestProductValidateNoBaseURL test validation method for a product entity
func TestProductValidateNoBaseURL(t *testing.T) {
	p, err := entity.NewProduct("Test Product", "")

	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, err, entity.ErrInvalidEntity)
}
