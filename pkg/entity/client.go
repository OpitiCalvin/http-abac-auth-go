package entity

import "time"

type Client struct {
	ID         int64     `json:"id"`
	ClientName string    `json:"client_name"`
	PartnerID  int64     `json:"partner_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Products   []int64   `json:"products"`
}

// NewClient create a new client
func NewClient(clientName string, partnerID int64) (*Client, error) {
	c := &Client{
		ClientName: clientName,
		PartnerID:  partnerID,
		CreatedAt:  time.Now(),
	}
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Validate validate a client record
func (c *Client) Validate() error {
	if c.ClientName == "" {
		return ErrInvalidEntity
	}
	return nil
}

// AddProduct add a product to a client
func (c *Client) AddProduct(id int64) error {
	_, err := c.GetProduct(id)
	if err == nil {
		return ErrClientAlreadySubscribedToProduct
	}

	c.Products = append(c.Products, id)
	return nil
}

// RemoveProduct remove a product from a client record
func (c *Client) RemoveProduct(id int64) error {
	for i, j := range c.Products {
		if j == id {
			c.Products = append(c.Products[:i], c.Products[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetProduct get a client-subscribed product
func (c *Client) GetProduct(id int64) (int64, error) {
	for _, v := range c.Products {
		if v == id {
			return id, nil
		}
	}
	return 0, ErrNotFound
}
