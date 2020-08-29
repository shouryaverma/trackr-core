package postgres

import (
	"errors"

	"github.com/amaraliou/trackr-v2/internal/model"
	"github.com/jinzhu/gorm"
)

// CreateUser ...
func (repo *Repository) CreateUser(user model.User) (*model.User, error) {

	db := repo.postgres.DB
	logger := repo.postgres.logger

	err := db.Create(&user).Error
	if err != nil {
		logger.Infof("Failed to create user in Postgres")
		return &model.User{}, err
	}

	return &user, nil
}

// AllUsers ...
func (repo *Repository) AllUsers() (*[]model.User, error) {

	db := repo.postgres.DB
	logger := repo.postgres.logger
	users := []model.User{}

	err := db.Model(&model.User{}).Limit(100).Find(&users).Error
	if err != nil {
		logger.Infof("Failed to get all users from Postgres")
		return &[]model.User{}, err
	}

	return &users, nil
}

// GetUser ...
func (repo *Repository) GetUser(id string) (*model.User, error) {

	db := repo.postgres.DB
	logger := repo.postgres.logger
	user := model.User{}

	err := db.Model(&model.User{}).Where("id = ?", id).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		logger.Infof("User not found in Postgres")
		return &model.User{}, errors.New("User not found")
	}

	if err != nil {
		logger.Infof("Failed to get the user from Postgres")
		return &model.User{}, err
	}

	return &user, nil
}

// GetUserByEmail ...
func (repo *Repository) GetUserByEmail(email string) (*model.User, error) {

	db := repo.postgres.DB
	logger := repo.postgres.logger
	user := model.User{}

	err := db.Model(&model.User{}).Where("email = ?", email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		logger.Infof("User not found in Postgres")
		return &model.User{}, errors.New("User not found")
	}

	if err != nil {
		logger.Infof("Failed to get the user from Postgres")
		return &model.User{}, err
	}

	return &user, nil
}

// UpdateUser ...
func (repo *Repository) UpdateUser(user model.User, id string) (*model.User, error) {

	db := repo.postgres.DB
	logger := repo.postgres.logger

	err := user.BeforeSave()
	if err != nil {
		return &model.User{}, err
	}

	err = db.Model(model.User{}).Updates(&user).Error
	if err != nil {
		logger.Infof("Failed to update the user in Postgres")
		return &model.User{}, err
	}

	return repo.GetUser(id)
}

// DeleteUser ...
func (repo *Repository) DeleteUser(id string) (int64, error) {

	db := repo.postgres.DB
	logger := repo.postgres.logger

	db = db.Unscoped().Model(&model.User{}).Where("id = ?", id).Take(&model.User{}).Delete(&model.User{})
	if gorm.IsRecordNotFoundError(db.Error) {
		logger.Infof("Failed to get the user from Postgres")
		return 0, errors.New("User not found")
	}

	if db.Error != nil {
		logger.Infof("Failed to delete the user from Postgres")
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
