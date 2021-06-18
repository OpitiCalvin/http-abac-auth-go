package user

import (
	"errors"
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/client"
)

// Service create user usecase service
type UserService struct {
	repo          Repository
	clientService client.UseCase
}

// NewUserService create new user usecase service
func NewUserService(r Repository, clientService client.UseCase) *UserService {
	return &UserService{
		repo:          r,
		clientService: clientService,
	}
}

// CreateUser create a user
func (s *UserService) CreateUser(email, username, password string, clientID int64) error {
	u, err := entity.NewUser(email, username, password, clientID)
	if err != nil {
		return err
	}

	// validate client id entry
	_, err = s.clientService.GetClient(u.ClientID)
	if err != nil {
		// return err
		return errors.New("no client with id provided")
	}

	return s.repo.Create(u)
}

// GetUser get a user record
func (s *UserService) GetUser(id int64) (*entity.User, error) {
	u, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, entity.ErrNotFound
	}

	return u, nil
}

// ListUsers list user records
func (s *UserService) ListUsers() ([]*entity.User, error) {
	users, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, entity.ErrNotFound
	}

	return users, nil
}

// UpdateUser update a user record
func (s *UserService) UpdateUser(e *entity.User) error {
	// validate client id entry
	_, err := s.clientService.GetClient(e.ClientID)
	if err != nil {
		return err
	}

	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// DeleteUser delete a user record
func (s *UserService) DeleteUser(id int64) error {
	_, err := s.GetUser(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
