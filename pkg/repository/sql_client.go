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
func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{
		db: db,
	}
}

// Create create a client record in a database table
func (r *ClientDB) Create(e *entity.Client) error {
	stmt, err := r.db.Prepare(`
		insert into client (client_name, partner_id, created_at)
		values(?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ClientName, e.PartnerID, e.CreatedAt)
	if err != nil {
		return err
	}

	// err = stmt.Close()
	// if err != nil {
	// 	return
	// }

	return nil
}

// List retrieves a list of client records
func (r *ClientDB) List() ([]*entity.Client, error) {
	stmt, err := r.db.Prepare(`
		select id, client_name, partner_id, created_at, updated_at
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
		err = rows.Scan(&c.ID, &c.ClientName, &c.PartnerID, &c.CreatedAt, &c.UpdatedAt)
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
	stmt, err := r.db.Prepare(`select id, client_name, partner_id, created_at, updated_at from client where id = ?`)
	if err != nil {
		return nil, err
	}

	var c entity.Client
	// rows, err := stmt.Query(id)
	// if err != nil {
	// 	return nil, err
	// }

	// for rows.Next() {
	// err = rows.Scan(&c.ID, &c.ClientName, &c.PartnerID, &c.CreatedAt, &c.UpdatedAt)
	// }

	row := stmt.QueryRow(id)
	err = row.Scan(&c.ID, &c.ClientName, &c.PartnerID, &c.CreatedAt, &c.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, err
	}

	stmt, err = r.db.Prepare(`select product_id from client_product where client_id = ?`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i int64
		err = rows.Scan(&i)
		c.Products = append(c.Products, i)
	}

	return &c, nil
}

// Update update a client record
func (r *ClientDB) Update(e *entity.Client) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update client set client_name = ?, partner_id = ?, updated_at = ? where id = ?",
		e.ClientName, e.PartnerID, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("delete from client_product where client_id = ?", e.ID)
	if err != nil {
		return err
	}

	for _, pID := range e.Products {
		_, err := r.db.Exec("insert into client_product values(?,?,?)", e.ID, pID, time.Now().Format("2006-01-02"))
		if err != nil {
			return err
		}
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
