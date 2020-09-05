package storage

import "github.com/amaraliou/trackr-core/internal/model"

// PostgresInterface ...
type PostgresInterface interface {
	CreateUser(model.User) (*model.User, error)
	GetUser(string) (*model.User, error)
	GetUserByEmail(string) (*model.User, error)
	UpdateUser(model.User, string) (*model.User, error)
	DeleteUser(string) (int64, error)
	AllUsers() (*[]model.User, error)

	CreateApplication(model.Application) (*model.Application, error)
	GetApplication(string) (*model.Application, error)
	UpdateApplication(model.Application, string) (*model.Application, error)
	DeleteApplication(string) (int64, error)
	AllApplications() (*[]model.Application, error)
}
