package client

import "github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

// Reader client reader interface
type Reader interface {
	Get(id int64) (*entity.Client, error)
	List() ([]*entity.Client, error)
}

// Writer client writer interface
type Writer interface {
	Create(e *entity.Client) (int64, error)
	Update(e *entity.Client) error
	Delete(id int64) error
}

// Repository client repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase client usecase interface
type UseCase interface {
	ListClients() ([]*entity.Client, error)
	GetClient(id int64) (*entity.Client, error)
	CreateClient(clientName string, partner_id int64) (int64, error)
	DeleteClient(id int64) error
	UpdateClient(e *entity.Client) error
}
