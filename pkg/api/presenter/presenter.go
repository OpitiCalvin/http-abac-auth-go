package presenter

import "time"

// Product
type Product struct {
	ID          int64     `json:"id"`
	ProductName string    `json:"product_name"`
	BaseURL     string    `json:"base_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProductCreated product presentation after creation
type ProductCreated struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	BaseURL     string `json:"base_url"`
}

// Partner data
type Partner struct {
	ID          int64     `json:"id"`
	PartnerName string    `json:"partner_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PartnerCreated partner presentation struct after record creation
type PartnerCreated struct {
	ID          int64  `json:"id"`
	PartnerName string `json:"partner_name"`
}

// Client data
type Client struct {
	ID         int64     `json:"id"`
	ClientName string    `json:"client_name"`
	Products   []int64   `json:"products"`
	PartnerID  int64     `json:"partner_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ClientCreated client presentation struct after record creation
type ClientCreated struct {
	ID         int64   `json:"id"`
	ClientName string  `json:"client_name"`
	Products   []int64 `json:"products"`
	PartnerID  int64   `json:"partner_id"`
}

// User data
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	ClientID  int64     `json:"client_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
