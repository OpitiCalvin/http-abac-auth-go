package entity

import "time"

type Client struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Products  []int     `json:"products"`
	PartnerID int64     `json:"partner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// UpdatedBy int       `json:"updated_by"`
}

// NewClient create a new client
func NewClient(name string) (*Client, error) {
	return &Client{
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}
