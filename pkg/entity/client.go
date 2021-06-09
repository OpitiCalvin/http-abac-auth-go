package entity

import "time"

type Client struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Products  []int64   `json:"products"`
	PartnerID int64     `json:"partner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// UpdatedBy int       `json:"updated_by"`
}

// NewClient create a new client
func NewClient(name string, products []int64, partnerID int64) (*Client, error) {
	c := &Client{
		Name:      name,
		Products:  products,
		PartnerID: partnerID,
		CreatedAt: time.Now(),
	}
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Validate validate a client record
func (c *Client) Validate() error {
	if c.Name == "" {
		return ErrInvalidEntity
	}
	return nil
}
