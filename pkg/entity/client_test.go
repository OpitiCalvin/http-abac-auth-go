package entity_test

import (
	"testing"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/stretchr/testify/assert"
)

// TestNewClient test creation of new client entity
func TestNewClient(t *testing.T) {
	c, err := entity.NewClient("Test Client")
	assert.Nil(t, err)
	assert.Equal(t, c.ClientName, "Test Client")
}

// TestNewClientValidateNoName
func TestNewClientValidateNoName(t *testing.T) {
	_, err := entity.NewClient("")

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrInvalidEntity)
}
