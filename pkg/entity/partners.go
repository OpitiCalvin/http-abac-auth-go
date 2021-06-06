package entity

import (
	"time"
)

type Partner struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPartner create a new partner record
func NewPartner(name string) (*Partner, error) {
	p := &Partner{
		Name:      name,
		CreatedAt: time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

// Validate validate a partner record
func (p *Partner) Validate() error {
	if p.Name == "" {
		// return errors.New("partner name is required")
		return ErrInvalidEntity
	}
	return nil
}
