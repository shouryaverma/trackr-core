package postgres

import (
	"errors"

	"github.com/amaraliou/trackr-core/internal/model"
)

// CreateApplication ...
func (repo *Repository) CreateApplication(application model.Application) (*model.Application, error) {

	db := repo.postgres.DB
	logger := repo.postgres.logger

	if application.UserID.String() == "00000000-0000-0000-0000-000000000000" {
		logger.Infof("Failed to create application in Postgres: User ID not given")
		return &model.Application{}, errors.New("Invalid User ID")
	}

	_, err := repo.GetUser(application.UserID.String())
	if err != nil {
		logger.Infof("Failed to create application in Postgres: User not found")
		return &model.Application{}, errors.New("User doesn't exist, can't create account")
	}

	err = db.Create(&application).Error
	if err != nil {
		logger.Warnf("Failed to create application in Postgres: %s", err.Error())
		return &model.Application{}, err
	}

	return &application, nil
}

// AllApplications ...
func (repo *Repository) AllApplications() (*[]model.Application, error) {
	return &[]model.Application{}, nil
}

// GetApplication ...
func (repo *Repository) GetApplication(id string) (*model.Application, error) {
	return &model.Application{}, nil
}

// UpdateApplication
func (repo *Repository) UpdateApplication(application model.Application, id string) (*model.Application, error) {
	return &model.Application{}, nil
}

// DeleteApplication
func (repo *Repository) DeleteApplication(id string) (int64, error) {
	return int64(1), nil
}
