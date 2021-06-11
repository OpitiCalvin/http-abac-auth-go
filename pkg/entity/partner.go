package entity

import (
	"time"
)

type Partner struct {
	ID          int       `json:"id"`
	PartnerName string    `json:"partner_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewPartner create a new partner record
func NewPartner(partnerName string) (*Partner, error) {
	p := &Partner{
		PartnerName: partnerName,
		CreatedAt:   time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Validate validate a partner record
func (p *Partner) Validate() error {
	if p.PartnerName == "" {
		// return errors.New("partner name is required")
		return ErrInvalidEntity
	}
	return nil
}
