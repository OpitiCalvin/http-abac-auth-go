package client

import (
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/partner"
)

// Service client usecase service
type ClientService struct {
	repo        Repository
	partnerRepo partner.Repository
}

// NewClientService new client usercase service
func NewClientService(r Repository, partnerRepo partner.Repository) *ClientService {
	return &ClientService{
		repo:        r,
		partnerRepo: partnerRepo,
	}
}

// CreateClient create a client
func (s *ClientService) CreateClient(clientName string, partnerID int64) (int64, error) {
	c, err := entity.NewClient(clientName, partnerID)
	if err != nil {
		return 0, err
	}

	// // validate product entries
	// for _, prodID := range c.Products {
	// 	_, err := s.productRepo.Get(prodID)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// }

	// validate partner id information
	_, err = s.partnerRepo.Get(c.PartnerID)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(c)
}

// GetClient get a client record
func (s *ClientService) GetClient(id int64) (*entity.Client, error) {
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
func (s *ClientService) ListClients() ([]*entity.Client, error) {
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
func (s *ClientService) UpdateClient(e *entity.Client) error {
	// validate partner id information
	_, err := s.partnerRepo.Get(e.PartnerID)
	if err != nil {
		return err
	}

	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// DeleteClient delete a client record
func (s *ClientService) DeleteClient(id int64) error {
	_, err := s.GetClient(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
