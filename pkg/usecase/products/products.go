package usecases

import "github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

type ProductUseCase interface {
	Fetch(cursor string, num int) ([]*entity.Product, error)
	FetchByID(id int) (*entity.Product, error)
	Update(p *entity.Product) (*entity.Product, error)
	Store(p *entity.Product) (int, error)
	Delete(id int) (bool, error)
}
