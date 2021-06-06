package partner

import (
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
)

// Service partber usercase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreatePartner create a partner
func (s *Service) CreatePartner(name string) (int, error) {
	p, err := entity.NewPartner(name)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(p)
}

// GerPartner get a partner record
func (s *Service) GetPartner(id int) (*entity.Partner, error) {
	p, err := s.repo.Get(id)
	if p == nil {
		// return nil, fmt.Errorf("no partner record found with id %i", id)
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}
	return p, nil
}

// ListPartners lists partner records
func (s *Service) ListPartners() ([]*entity.Partner, error) {
	partners, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(partners) == 0 {
		// return nil, errors.New("no partner records found")
		return nil, entity.ErrNotFound
	}

	return partners, nil
}

// DeletePartner delete a partner record
func (s *Service) DeletePartner(id int) error {
	_, err := s.GetPartner(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// UpdatePartner update a partner record
func (s *Service) UpdatePartner(e *entity.Partner) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
