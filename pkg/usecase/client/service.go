package client

import (
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/partner"
)

// Service client usecase service
type ClientService struct {
	repo Repository
	// partnerRepo partner.Repository
	partnerService partner.UseCase
}

// NewClientService new client usercase service
func NewClientService(r Repository, partnerService partner.UseCase) *ClientService {
	return &ClientService{
		repo:           r,
		partnerService: partnerService,
	}
}

// CreateClient create a client
func (s *ClientService) CreateClient(clientName string) (int64, error) {
	c, err := entity.NewClient(clientName)
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
	err := e.Validate()
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

// LinkToPartner link a client record to a partner record
func (c *ClientService) LinkToPartner(e *entity.Client, partnerID int64) error {
	_, err := c.partnerService.GetPartner(partnerID)
	if err != nil {
		return entity.ErrNotFound
	}

	err = e.AddPartner(partnerID)
	if err != nil {
		return err
	}

	err = c.UpdateClient(e)
	if err != nil {
		return err
	}

	return nil
}

// UnlinkFromPartner unlink a client record from a partner
func (c *ClientService) UnlinkFromPartner(e *entity.Client) error {
	if e.PartnerID == 0 {
		return nil
	}
	e.PartnerID = 0
	err := c.UpdateClient(e)
	if err != nil {
		return err
	}
	return nil
}
