package postgres

import (
	"errors"

	"github.com/amaraliou/trackr-core/internal/model"
	"github.com/jinzhu/gorm"
)

// CreateApplication ...
func (repo *Repository) CreateApplication(application model.Application) (*model.Application, error) {

	db := repo.postgres.DB
	logger := repo.postgres.logger

	if application.UserID.String() == "00000000-0000-0000-0000-000000000000" {
		logger.Infof("Failed to create application in Postgres: Application ID not given")
		return &model.Application{}, errors.New("Invalid Application ID")
	}

	_, err := repo.GetUser(application.UserID.String())
	if err != nil {
		logger.Infof("Failed to create application in Postgres: Application not found")
		return &model.Application{}, errors.New("User doesn't exist, can't create application")
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

	db := repo.postgres.DB
	logger := repo.postgres.logger
	applications := []model.Application{}

	err := db.Model(&model.Application{}).Limit(100).Find(&applications).Error
	if err != nil {
		logger.Warnf("Failed to retrieve applications in Postgres: %s", err.Error())
		return &[]model.Application{}, err
	}

	return &applications, nil
}

// AllUserApplications ...
func (repo *Repository) AllUserApplications(ApplicationID string) (*[]model.Application, error) {
	return &[]model.Application{}, nil
}

// GetApplication ...
func (repo *Repository) GetApplication(id string) (*model.Application, error) {

	db := repo.postgres.DB
	logger := repo.postgres.logger
	application := model.Application{}

	err := db.Model(&model.Application{}).Where("id = ?", id).Take(&application).Error
	if gorm.IsRecordNotFoundError(err) {
		logger.Infof("Application not found in Postgres")
		return &model.Application{}, errors.New("Application not found")
	}

	if err != nil {
		logger.Infof("Failed to get the application from Postgres")
		return &model.Application{}, err
	}

	return &application, nil
}

// UpdateApplication ...
func (repo *Repository) UpdateApplication(application model.Application, id string) (*model.Application, error) {
	return &model.Application{}, nil
}

// DeleteApplication ...
func (repo *Repository) DeleteApplication(id string) (int64, error) {
	return int64(1), nil
}
