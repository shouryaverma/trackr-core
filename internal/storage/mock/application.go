package mock

import (
	"errors"

	"github.com/amaraliou/trackr-core/internal/model"
)

// CreateApplication ...
func (repo *Repository) CreateApplication(application model.Application) (*model.Application, error) {

	returnObject := repo.ReturnObject.(*model.Application)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}

// GetApplication ...
func (repo *Repository) GetApplication(id string) (*model.Application, error) {

	returnObject := repo.ReturnObject.(*model.Application)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}

// UpdateApplication ...
func (repo *Repository) UpdateApplication(application model.Application, id string) (*model.Application, error) {

	returnObject := repo.ReturnObject.(*model.Application)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}

// DeleteApplication ...
func (repo *Repository) DeleteApplication(id string) (int64, error) {

	returnObject := repo.ReturnObject.(int64)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}

// AllApplications ...
func (repo *Repository) AllApplications() (*[]model.Application, error) {

	returnObject := repo.ReturnObject.(*[]model.Application)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}
