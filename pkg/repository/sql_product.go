package repository

import (
	"database/sql"
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
)

type ProductDB struct {
	db *sql.DB
}

// NewProductDB create new product repository
func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{
		db: db,
	}
}

// Create create a product record in a database table
func (r *ProductDB) Create(e *entity.Product) (int64, error) {
	stmt, err := r.db.Prepare(`
		insert into product (name, base_url, created_at)
		values(?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.BaseURL, e.CreatedAt)
	if err != nil {
		return 0, err
	}

	// err = stmt.Close()
	// if err != nil {
	// 	return
	// }

	// get last inserted id
	lid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lid, nil
}

// List retrieves a list of product records
func (r *ProductDB) List() ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(`select id, name, base_url, created_at, updated_at from product`)
	if err != nil {
		return nil, err
	}

	var products []*entity.Product
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var p entity.Product
		err = rows.Scan(&p.ID, &p.Name, &p.BaseURL, &p.CreatedAt, &p.UpdatedAt)
		// TODO: -handle row scan with null datetime values (updated_at column)
		// if err != nil {
		// 	return nil, err
		// }
		products = append(products, &p)
	}
	return products, nil
}

// Get retrieve a product record using its unique id
func (r *ProductDB) Get(id int64) (*entity.Product, error) {
	stmt, err := r.db.Prepare(`
		select id, name, base_url, created_at, updated_at
		from product
		where id = ?
	`)
	if err != nil {
		return nil, err
	}

	var product entity.Product
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.BaseURL, &product.CreatedAt, &product.UpdatedAt)
	}
	return &product, nil
}

// Update update a product record
func (r *ProductDB) Update(e *entity.Product) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update product set name = ?, base_url = ?, updated_at = ? where id = ?",
		e.Name, e.BaseURL, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete delete a product record
func (r *ProductDB) Delete(id int64) error {
	_, err := r.db.Exec("delete from product where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
