package product

import (
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
)

// Service product usecase
type Service struct {
	repo Repository
}

// NewService create new product service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateProduct create a product
func (s *Service) CreateProduct(name, baseURL string) (int64, error) {
	p, err := entity.NewProduct(name, baseURL)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(p)
}

// GerProduct get a product record
func (s *Service) GetProduct(id int64) (*entity.Product, error) {
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
func (s *Service) ListProducts() ([]*entity.Product, error) {
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
func (s *Service) DeleteProduct(id int64) error {
	_, err := s.GetProduct(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// UpdateProduct update a product record
func (s *Service) UpdateProduct(e *entity.Product) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
