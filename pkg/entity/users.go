package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	ClientID  int64     `json:"client_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// validate validate user data
func (u *User) validate() error {
	if u.Email == "" {
		return errors.New("email address is required")
	}

	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

// create password hash
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	return string(hash), nil
}

// validate password
func (u *User) checkPassword(password string) (isPasswordValid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}

	return true
}

// NewUser create a new user
func NewUser(email, username, password string, clientId int) (*User, error) {
	// TODO: get Client ID to add to this new user record

	u := &User{
		Email:     email,
		Username:  username,
		CreatedAt: time.Now(),
	}

	pwd, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd

	err = u.validate()
	if err != nil {
		return nil, err
	}
	return u, nil
}
