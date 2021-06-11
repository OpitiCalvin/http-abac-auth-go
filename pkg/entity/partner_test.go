package entity_test

import (
	"testing"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/stretchr/testify/assert"
)

// TestNewPartner test creation of new partner object
func TestNewPartner(t *testing.T) {
	p, err := entity.NewPartner("Test Partner")

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.PartnerName, "Test Partner")
}

// TestPartnerValidate
func TestPartnerValidate(t *testing.T) {
	_, err := entity.NewPartner("")

	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrInvalidEntity)
}
