package product

import "github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

// Reader interface
type Reader interface {
	Get(id int64) (*entity.Product, error)
	List() ([]*entity.Product, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.Product) error
	Update(e *entity.Product) error
	Delete(id int64) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	ListProducts() ([]*entity.Product, error)
	GetProduct(id int64) (*entity.Product, error)
	CreateProduct(productName, baseURL string) error
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id int64) error
}
