package repository

import (
	"database/sql"
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
)

// PartnerDB postgresql repo
type PartnerDB struct {
	db *sql.DB
}

// NewPartnerDB create new repository
func NewPartnerDB(db *sql.DB) *PartnerDB {
	return &PartnerDB{
		db: db,
	}
}

// Create create a partner record
func (r *PartnerDB) Create(e *entity.Partner) (int64, error) {
	stmt, err := r.db.Prepare(`insert into partner (partner_name, created_at) values(?,?)`)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(
		e.PartnerName,
		time.Now().Format("2006-01-02"),
	)

	if err != nil {
		return 0, err
	}
	err = stmt.Close()
	if err != nil {
		return 0, err
	}

	// TODO: switch and instead return id of new record created
	lid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lid, nil
}

// Get get a partner record
func (r *PartnerDB) Get(id int64) (*entity.Partner, error) {
	stmt, err := r.db.Prepare(`select id, partner_name, created_at, updated_at from partner where id =?`)
	if err != nil {
		return nil, err
	}
	var p entity.Partner
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.PartnerName, &p.CreatedAt, &p.UpdatedAt)
	}
	return &p, nil
}

// Update update a partner record
func (r *PartnerDB) Update(e *entity.Partner) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update partner set partner_name = ?, updated_at = ? where id = ?",
		e.PartnerName, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete delete a partner record
func (r *PartnerDB) Delete(id int64) error {
	_, err := r.db.Exec("delete from partner where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// List get list of partner records
func (r *PartnerDB) List() ([]*entity.Partner, error) {
	stmt, err := r.db.Prepare(`select id, partner_name, created_at, updated_at from partner`)
	if err != nil {
		return nil, err
	}

	var partners []*entity.Partner
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Partner
		err = rows.Scan(&b.ID, &b.PartnerName, &b.CreatedAt, &b.UpdatedAt)
		// TODO - handle row scans with null datetime values

		// if err != nil {
		// 	return nil, err
		// }
		partners = append(partners, &b)
	}
	return partners, nil
}
