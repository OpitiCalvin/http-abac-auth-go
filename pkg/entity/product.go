package entity

import "time"

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BaseURL   string    `json:"base_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewProduct create a new product record
func NewProduct(name, baseURL string) (*Product, error) {
	p := &Product{
		Name:      name,
		BaseURL:   baseURL,
		CreatedAt: time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Validate validate a product record
func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrInvalidEntity
	}

	if p.BaseURL == "" {
		return ErrInvalidEntity
	}

	return nil
}
