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

// Validate validate user data
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email address is required")
	}

	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	if u.ClientID == int64(0) {
		return errors.New("a valid client id is required")
	}

	return nil
}

// hashPassword create password hash
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	return string(hash), nil
}

// CheckPassword validate password
func (u *User) CheckPassword(password string) (isPasswordValid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}

	return true
}

// NewUser create a new user
func NewUser(email, username, password string, clientId int64) (*User, error) {
	// TODO: get Client ID to add to this new user record

	u := &User{
		Email:     email,
		Username:  username,
		Password:  password,
		ClientID:  clientId,
		CreatedAt: time.Now(),
	}

	err := u.Validate()
	if err != nil {
		return nil, err
	}

	pwd, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	u.Password = pwd

	return u, nil
}
