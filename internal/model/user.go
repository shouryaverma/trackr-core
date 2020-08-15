package model

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Base -> Struct to substitute gorm.Model when I want a UUID
type Base struct {
	ID        uuid.UUID  `gorm:"primary_key" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// User ...
type User struct {
	Base
	Email      string `json:"email" gorm:"unique;not null"`
	Password   string `json:"password"`
	IsVerified bool   `json:"is_verified"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}

// Hash -> Generate hash for given password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword -> Verify a password given it's hash
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// BeforeSave will check hashes for passwords
func (user *User) BeforeSave() error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// Validate user
func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if user.Email == "" {
			return errors.New("Required Email")
		}

		if user.Password == "" {
			return errors.New("Required Password")
		}

		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}

		if user.FirstName == "" {
			return errors.New("Required First Name")
		}

		if user.LastName == "" {
			return errors.New("Required Last Name")
		}

		return nil

	case "login":
		if user.Email == "" {
			return errors.New("Required Email")
		}

		if user.Password == "" {
			return errors.New("Required Password")
		}

		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	default:
		return nil
	}
}
