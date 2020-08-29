package storage

import "github.com/amaraliou/trackr-v2/internal/model"

// PostgresInterface ...
type PostgresInterface interface {
	CreateUser(model.User) (*model.User, error)
	GetUser(string) (*model.User, error)
	GetUserByEmail(string) (*model.User, error)
	UpdateUser(model.User, string) (*model.User, error)
	DeleteUser(string) (int64, error)
	AllUsers() (*[]model.User, error)
}
