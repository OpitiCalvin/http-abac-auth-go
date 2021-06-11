package product

import (
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
)

// Service product usecase
type ProductService struct {
	repo Repository
}

// NewProductService create new product service
func NewProductService(r Repository) *ProductService {
	return &ProductService{
		repo: r,
	}
}

// CreateProduct create a product
func (s *ProductService) CreateProduct(productName, baseURL string) (int64, error) {
	p, err := entity.NewProduct(productName, baseURL)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(p)
}

// GerProduct get a product record
func (s *ProductService) GetProduct(id int64) (*entity.Product, error) {
	p, err := s.repo.Get(id)
	if p == nil {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}
	return p, nil
}

// ListProducts lists product records
func (s *ProductService) ListProducts() ([]*entity.Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entity.ErrNotFound
	}

	return products, nil
}

// DeleteProduct delete a product record
func (s *ProductService) DeleteProduct(id int64) error {
	_, err := s.GetProduct(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// UpdateProduct update a product record
func (s *ProductService) UpdateProduct(e *entity.Product) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
