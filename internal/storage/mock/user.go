package mock

import (
	"errors"

	"github.com/amaraliou/trackr-core/internal/model"
)

// CreateUser ...
func (repo *Repository) CreateUser(user model.User) (*model.User, error) {

	returnObject := repo.ReturnObject.(*model.User)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}

// GetUser ...
func (repo *Repository) GetUser(id string) (*model.User, error) {

	returnObject := repo.ReturnObject.(*model.User)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}

// GetUserByEmail ...
func (repo *Repository) GetUserByEmail(email string) (*model.User, error) {

	returnObject := repo.ReturnObject.(*model.User)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}

// UpdateUser ...
func (repo *Repository) UpdateUser(user model.User, id string) (*model.User, error) {

	returnObject := repo.ReturnObject.(*model.User)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}

// DeleteUser ...
func (repo *Repository) DeleteUser(id string) (int64, error) {

	returnObject := repo.ReturnObject.(int64)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}

// AllUsers ...
func (repo *Repository) AllUsers() (*[]model.User, error) {

	returnObject := repo.ReturnObject.(*[]model.User)

	if repo.IsError {
		return returnObject, errors.New(repo.ErrorMessage)
	}

	return returnObject, nil
}
