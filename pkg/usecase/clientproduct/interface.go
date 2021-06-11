package clientproduct

import "github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

type UseCase interface {
	Subscribe(c *entity.Client, p *entity.Product) error
	Unsubscribe(c *entity.Client, p *entity.Product) error
}
