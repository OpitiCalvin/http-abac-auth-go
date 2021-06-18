package partner

import "github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

// Reader interface
type Reader interface {
	Get(id int64) (*entity.Partner, error)
	List() ([]*entity.Partner, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.Partner) error
	Update(e *entity.Partner) error
	Delete(id int64) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	ListPartners() ([]*entity.Partner, error)
	GetPartner(id int64) (*entity.Partner, error)
	CreatePartner(partnerName string) error
	UpdatePartner(e *entity.Partner) error
	DeletePartner(id int64) error
}
