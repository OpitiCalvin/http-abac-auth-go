package product

import (
	"testing"
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureProduct() *entity.Product {
	return &entity.Product{
		ID:          int64(1),
		ProductName: "Product One",
		BaseURL:     "http://product-one.com",
		CreatedAt:   time.Now(),
	}
}

func TestCreateProduct(t *testing.T) {
	repo := newInmem()
	m := NewProductService(repo)
	e := newFixtureProduct()
	err := m.CreateProduct(e.ProductName, e.BaseURL)
	assert.Nil(t, err)
	assert.False(t, e.CreatedAt.IsZero())
}

func TestListAndGetProduct(t *testing.T) {
	repo := newInmem()
	m := NewProductService(repo)
	e1 := newFixtureProduct()
	e2 := newFixtureProduct()
	e2.ProductName = "Product Two"
	e2.BaseURL = "http://product-two.com"

	_ = m.CreateProduct(e1.ProductName, e1.BaseURL)
	_ = m.CreateProduct(e2.ProductName, e2.BaseURL)

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListProducts()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		p, err := m.GetProduct(int64(1))
		assert.Nil(t, err)
		assert.Equal(t, e1.ProductName, p.ProductName)
	})
}

func TestUpdateProduct(t *testing.T) {
	repo := newInmem()
	m := NewProductService(repo)
	e := newFixtureProduct()
	err := m.CreateProduct(e.ProductName, e.BaseURL)
	assert.Nil(t, err)

	p, err := m.GetProduct(int64(1))
	assert.Nil(t, err)
	p.ProductName = "Product One Updated"
	err = m.UpdateProduct(p)
	assert.Nil(t, err)
	updated, err := m.GetProduct(p.ID)
	assert.Nil(t, err)
	assert.Equal(t, "Product One Updated", updated.ProductName)
}

func TestDeleteProduct(t *testing.T) {
	repo := newInmem()
	m := NewProductService(repo)
	e := newFixtureProduct()
	_ = m.CreateProduct(e.ProductName, e.BaseURL)

	err := m.DeleteProduct(int64(2))
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteProduct(int64(1))
	assert.Nil(t, err)
	_, err = m.GetProduct(int64(1))
	assert.Equal(t, entity.ErrNotFound, err)
}
