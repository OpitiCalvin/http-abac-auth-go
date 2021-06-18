package repository

import (
	"database/sql"
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
)

type UserDB struct {
	db *sql.DB
}

// NewUserDB create a new user repository
func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

// Create create a user record in a database table
func (r *UserDB) Create(e *entity.User) error {
	stmt, err := r.db.Prepare(`
		insert into user (email, username, password, client_id, created_at)
		values(?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Email, e.Username, e.Password, e.ClientID, e.CreatedAt)
	if err != nil {
		return err
	}

	// err = stmt.Close()
	// if err != nil {
	// 	return
	// }

	return nil
}

// List retrieves a list of user records
func (r *UserDB) List() ([]*entity.User, error) {
	stmt, err := r.db.Prepare(`
		select id, email, username, password, client_id, created_at, updated_at
		from user
	`)
	if err != nil {
		return nil, err
	}

	var users []*entity.User
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var u entity.User
		err = rows.Scan(&u.ID, &u.Email, &u.Username, &u.Password, &u.ClientID, &u.CreatedAt, &u.UpdatedAt)
		// TODO: -handle row scan with null datetime values (updated_at column)
		// if err != nil {
		// 	return nil, err
		// }
		users = append(users, &u)
	}
	return users, nil
}

// Get retrieve a user record using its unique id
func (r *UserDB) Get(id int64) (*entity.User, error) {
	stmt, err := r.db.Prepare(`
		select id, email, username, password, client_id, created_at, updated_at
		from user
		where id = ?
	`)
	if err != nil {
		return nil, err
	}

	var user entity.User
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.ClientID, &user.CreatedAt, &user.UpdatedAt)
	}
	return &user, nil
}

// Update update a user record
func (r *UserDB) Update(e *entity.User) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update user set email = ?, username = ?, password = ?, client_id = ?, updated_at = ? where id = ?",
		e.Email, e.Username, e.Password, e.ClientID, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete delete a user record
func (r *UserDB) Delete(id int64) error {
	_, err := r.db.Exec("delete from user where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
