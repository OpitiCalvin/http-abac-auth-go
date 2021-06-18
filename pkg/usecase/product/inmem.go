package product

import (
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
)

// inmem in memory repo
type inmem struct {
	m map[int64]*entity.Product
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[int64]*entity.Product{}
	return &inmem{
		m: m,
	}
}

func (r *inmem) getLastMapKey() int64 {
	keys := make([]int64, 0, len(r.m))
	for k := range r.m {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		return int64(0)
	}

	return keys[len(keys)-1]
}

// Create a Product
func (r *inmem) Create(e *entity.Product) error {
	nextID := r.getLastMapKey()
	nextID = nextID + 1
	e.ID = nextID

	r.m[nextID] = e
	return nil
}

// Get a product
func (r *inmem) Get(id int64) (*entity.Product, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update a product
func (r *inmem) Update(e *entity.Product) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// List products
func (r *inmem) List() ([]*entity.Product, error) {
	var d []*entity.Product
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete a Product
func (r *inmem) Delete(id int64) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
