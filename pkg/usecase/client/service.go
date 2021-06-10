package client

import (
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/partner"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/product"
)

// Service client usecase service
type Service struct {
	repo           Repository
	productService product.Service
	partnerService partner.Service
}

// NewService new client usercase service
func NewService(r Repository, prs product.Service, pts partner.Service) *Service {
	return &Service{
		repo:           r,
		productService: prs,
		partnerService: pts,
	}
}

// CreateClient create a client
func (s *Service) CreateClient(name string, products []int64, partnerID int64) (int64, error) {
	c, err := entity.NewClient(name, products, partnerID)
	if err != nil {
		return 0, err
	}

	// validate product entries
	for _, prodID := range c.Products {
		_, err := s.productService.GetProduct(prodID)
		if err != nil {
			return 0, err
		}
	}

	// validate partner id information
	_, err = s.partnerService.GetPartner(c.PartnerID)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(c)
}

// GetClient get a client record
func (s *Service) GetClient(id int64) (*entity.Client, error) {
	c, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, entity.ErrNotFound
	}
	return c, nil
}

// ListClients list client records
func (s *Service) ListClients() ([]*entity.Client, error) {
	clients, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(clients) == 0 {
		return nil, entity.ErrNotFound
	}
	return clients, nil
}

// UpdateClient update a client record
func (s *Service) UpdateClient(e *entity.Client) error {
	// validate product entries
	for _, prodID := range e.Products {
		_, err := s.productService.GetProduct(prodID)
		if err != nil {
			return err
		}
	}

	// validate partner id information
	_, err := s.partnerService.GetPartner(e.PartnerID)
	if err != nil {
		return err
	}

	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// DeleteClient delete a client record
func (s *Service) DeleteClient(id int64) error {
	_, err := s.GetClient(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
