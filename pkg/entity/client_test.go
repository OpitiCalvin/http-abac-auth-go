package entity_test

import (
	"testing"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/stretchr/testify/assert"
)

// TestNewClient test creation of new client entity
func TestNewClient(t *testing.T) {
	c, err := entity.NewClient("Test Client", []int64{1, 3}, int64(1))
	assert.Nil(t, err)
	assert.Equal(t, c.Name, "Test Client")
	assert.Equal(t, c.Products, []int64{1, 3})
	assert.Equal(t, c.PartnerID, int64(1))
}

// TestNewClientValidateNoName
func TestNewClientValidateNoName(t *testing.T) {
	_, err := entity.NewClient("", []int64{1, 3}, int64(1))

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrInvalidEntity)
}
