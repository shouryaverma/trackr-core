// +build integration

package postgres

import (
	"log"
	"testing"

	"github.com/amaraliou/trackr-core/internal/model"
	"gopkg.in/go-playground/assert.v1"
)

func TestAllApplications(t *testing.T) {

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	applications, err := seedApplications()
	if err != nil {
		log.Fatal(err)
	}

	retrievedApplications, err := pgRepo.AllApplications()
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(*applications), len(*retrievedApplications))
}

func TestGetApplication(t *testing.T) {

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	application, err := seedOneApplication()
	if err != nil {
		log.Fatal(err)
	}

	retrievedApplication, err := pgRepo.GetApplication(application.ID.String())
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, application.JobTitle, retrievedApplication.JobTitle)
	assert.Equal(t, application.Company, retrievedApplication.Company)
}

func TestCreateApplication(t *testing.T) {

	err := refreshEverything()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	newApplication := model.Application{
		JobTitle: "Software Engineer Intern",
		Company:  "GoCardless",
		UserID:   user.ID,
	}

	createdApplication, err := pgRepo.CreateApplication(newApplication)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, newApplication.JobTitle, createdApplication.JobTitle)
}
