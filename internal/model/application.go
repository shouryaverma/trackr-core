package model

import (
	"errors"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// Application ..
type Application struct {
	Base
	JobTitle    string    `json:"job_title"`
	Company     string    `json:"company"` // Add struct/table for company details later
	Description string    `json:"description"`
	JobPosting  string    `json:"job_url"`
	Location    string    `json:"location"`
	Status      int       `json:"status"` // Define iota
	Type        string    `json:"type"`
	User        User      `json:"-" gorm:"foreignkey:UserID"`
	UserID      uuid.UUID `json:"user_id" gorm:"user_id"`
}

// Validate ..
func (application *Application) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if application.JobTitle == "" {
			return errors.New("Required Job Title")
		}

		if application.Company == "" {
			return errors.New("Required Company")
		}

		if application.UserID.String() == "00000000-0000-0000-0000-000000000000" {
			return errors.New("Required User ID")
		}

		return nil

	default:
		return nil
	}
}
