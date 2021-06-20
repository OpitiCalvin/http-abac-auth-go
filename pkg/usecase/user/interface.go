package user

import "github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

// Reader user reader interface
type Reader interface {
	Get(id int64) (*entity.User, error)
	List() ([]*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

// Writer user writer interface
type Writer interface {
	Create(e *entity.User) error
	Update(e *entity.User) error
	Delete(id int64) error
}

// Repository user interface
type Repository interface {
	Reader
	Writer
}

// UseCase user usecase interface
type UseCase interface {
	ListUsers() ([]*entity.User, error)
	GetUser(id int64) (*entity.User, error)
	CreateUser(email, username, password string, clientID int64) error
	UpdateUser(e *entity.User) error
	DeleteUser(id int64) error
	LoginUser(username, password string) map[string]interface{}
}
