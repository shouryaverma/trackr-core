// +build integration

package postgres

import (
	"log"
	"testing"

	"github.com/amaraliou/trackr-core/internal/model"
	"gopkg.in/go-playground/assert.v1"
)

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
