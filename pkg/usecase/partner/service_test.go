package partner

import (
	"testing"
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func newFixturePartner() *entity.Partner {
	return &entity.Partner{
		ID:          int64(1),
		PartnerName: "Test Partner",
		CreatedAt:   time.Now(),
	}
}

func TestCreatePartner(t *testing.T) {
	repo := newInmem()
	m := NewPartnerService(repo)
	e := newFixturePartner()
	err := m.CreatePartner(e.PartnerName)
	assert.Nil(t, err)
	assert.False(t, e.CreatedAt.IsZero())
}

func TestListAndGetPartner(t *testing.T) {
	repo := newInmem()
	m := NewPartnerService(repo)
	e1 := newFixturePartner()
	e2 := newFixturePartner()
	e2.PartnerName = "Test Partner 2"

	_ = m.CreatePartner(e1.PartnerName)
	_ = m.CreatePartner(e2.PartnerName)

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListPartners()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		p, err := m.GetPartner(int64(1))
		assert.Nil(t, err)
		assert.Equal(t, e1.PartnerName, p.PartnerName)
	})
}

func TestUpdatePartner(t *testing.T) {
	repo := newInmem()
	m := NewPartnerService(repo)
	e := newFixturePartner()
	err := m.CreatePartner(e.PartnerName)
	assert.Nil(t, err)

	p, err := m.GetPartner(int64(1))
	assert.Nil(t, err)
	p.PartnerName = "Test Partner Updated"
	err = m.UpdatePartner(p)
	assert.Nil(t, err)
	updated, err := m.GetPartner(p.ID)
	assert.Nil(t, err)
	assert.Equal(t, "Test Partner Updated", updated.PartnerName)
}

func TestDeletePartner(t *testing.T) {
	repo := newInmem()
	m := NewPartnerService(repo)
	e := newFixturePartner()
	_ = m.CreatePartner(e.PartnerName)

	err := m.DeletePartner(int64(2))
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeletePartner(int64(1))
	assert.Nil(t, err)
	_, err = m.GetPartner(int64(1))
	assert.Equal(t, entity.ErrNotFound, err)
}
