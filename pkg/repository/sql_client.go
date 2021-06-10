package repository

import (
	"database/sql"
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
)

type ClientDB struct {
	db *sql.DB
}

// NewClientDB create new client repository
func NewClientDB(db *sql.DB) *PartnerDB {
	return &PartnerDB{
		db: db,
	}
}

// Create create a client record in a database table
func (r *ClientDB) Create(e *entity.Client) (int64, error) {
	stmt, err := r.db.Prepare(`
		insert into client (name, products, partner_id, created_at)
		values(?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Products, e.PartnerID, e.CreatedAt)
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

// List retrieves a list of client records
func (r *ClientDB) List() ([]*entity.Client, error) {
	stmt, err := r.db.Prepare(`
		select id, name, products, partner_id, created_at, updated_at
		from client
	`)
	if err != nil {
		return nil, err
	}

	var clients []*entity.Client
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var c entity.Client
		err = rows.Scan(&c.ID, &c.Name, &c.Products, &c.PartnerID, &c.CreatedAt, &c.UpdatedAt)
		// TODO: -handle row scan with null datetime values (updated_at column)
		// if err != nil {
		// 	return nil, err
		// }
		clients = append(clients, &c)
	}
	return clients, nil
}

// Get retrieve a client record using its unique id
func (r *ClientDB) Get(id int64) (*entity.Client, error) {
	stmt, err := r.db.Prepare(`
		select id, name, products, parner_id, created_at, updated_at
		from client
		where id = ?
	`)
	if err != nil {
		return nil, err
	}

	var c entity.Client
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		err = rows.Scan(&c.ID, &c.Name, &c.Products, &c.PartnerID, &c.CreatedAt, &c.UpdatedAt)
	}
	return &c, nil
}

// Update update a client record
func (r *ClientDB) Update(e *entity.Client) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update client set name = ?, products = ?, partner_id = ?, updated_at = ? where id = ?",
		e.Name, e.Products, e.PartnerID, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete delete a client record
func (r *ClientDB) Delete(id int64) error {
	_, err := r.db.Exec("delete from client where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
