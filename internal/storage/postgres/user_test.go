// +build integration

package postgres

import (
	"log"
	"testing"

	"github.com/amaraliou/trackr-core/internal/model"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/assert.v1"
)

func TestAllUsers(t *testing.T) {

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	users, err := seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	retrievedUsers, err := pgRepo.AllUsers()
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(*users), len(*retrievedUsers))
}

func TestGetUser(t *testing.T) {

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	retrievedUser, err := pgRepo.GetUser(user.ID.String())
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, user.Email, retrievedUser.Email)
}

func TestNonExistentUser(t *testing.T) {

	randomUUID := uuid.NewV4()
	randomEmail := "shbachfad@gjsda.com"

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	_, err = pgRepo.GetUser(randomUUID.String())
	assert.Equal(t, err.Error(), "User not found")

	_, err = pgRepo.GetUserByEmail(randomEmail)
	assert.Equal(t, err.Error(), "User not found")

	_, err = pgRepo.DeleteUser(randomUUID.String())
	assert.Equal(t, err.Error(), "User not found")
}

func TestGetUserByEmail(t *testing.T) {

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	retrievedUser, err := pgRepo.GetUserByEmail(user.Email)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, user.ID, retrievedUser.ID)
}

func TestCreateUser(t *testing.T) {

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	newUser := model.User{
		Email:     "johndoe@gmail.com",
		Password:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
	}

	createdUser, err := pgRepo.CreateUser(newUser)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, newUser.Email, createdUser.Email)
}

func TestUpdateUser(t *testing.T) {

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	userUpdate := model.User{
		FirstName: "Mario",
	}

	updatedUser, err := pgRepo.UpdateUser(userUpdate, user.ID.String())
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, updatedUser.FirstName, "Mario")
}

func TestDeleteUser(t *testing.T) {

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	isDeleted, err := pgRepo.DeleteUser(user.ID.String())
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, isDeleted, int64(1))
}
