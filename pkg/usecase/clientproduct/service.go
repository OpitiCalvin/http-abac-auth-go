package clientproduct

import (
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/client"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/product"
)

// ClientProductService client-product usecase
type ClientProductService struct {
	clientService  client.UseCase
	productService product.UseCase
}

// NewClientProductService create a new use case
func NewClientProductService(cs client.UseCase, ps product.UseCase) *ClientProductService {
	return &ClientProductService{
		clientService:  cs,
		productService: ps,
	}
}

// Subscribe subscribe or add a product to a client record
func (s *ClientProductService) Subscribe(c *entity.Client, p *entity.Product) error {
	c, err := s.clientService.GetClient(c.ID)
	if err != nil {
		return err
	}

	p, err = s.productService.GetProduct(p.ID)
	if err != nil {
		return err
	}

	err = c.AddProduct(int64(p.ID))
	if err != nil {
		return err
	}

	err = s.clientService.UpdateClient(c)
	if err != nil {
		return err
	}

	return nil
}

// Unsubscribe unsubscribe or remove a product from a client record
func (s *ClientProductService) Unsubscribe(c *entity.Client, p *entity.Product) error {
	c, err := s.clientService.GetClient(c.ID)
	if err != nil {
		return err
	}

	p, err = s.productService.GetProduct(p.ID)
	if err != nil {
		return err
	}

	err = c.RemoveProduct(int64(p.ID))
	if err != nil {
		return err
	}

	err = s.clientService.UpdateClient(c)
	if err != nil {
		return err
	}

	return nil
}
