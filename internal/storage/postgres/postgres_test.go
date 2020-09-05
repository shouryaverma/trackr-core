// +build integration

package postgres

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/amaraliou/trackr-core/internal/model"
	"github.com/amaraliou/trackr-core/pkg/logger"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var pgConn = Connection{}
var pgRepo = Repository{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../../.env"))
	if err != nil {
		fmt.Printf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_USER"),
		os.Getenv("TEST_DB_NAME"),
		os.Getenv("TEST_DB_PASSWORD"),
	)
	pgConn.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		log.Fatal(err)
	}

	logConfig := logger.Config{
		EnableConsole:     false,
		ConsoleLevel:      logger.Debug,
		ConsoleJSONFormat: true,
		EnableFile:        false,
		FileLevel:         logger.Info,
		FileJSONFormat:    true,
		FileLocation:      "log.log",
	}

	pgConn.logger, err = logger.NewZapLogger(logConfig)
	if err != nil {
		log.Fatal(err)
	}

	pgRepo.postgres = &pgConn

	if err != nil {
		fmt.Print("Cannot connect to Postgres database\n")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Print("We are connected to the Postgres database\n")
	}
}

func refreshEverything() error {

	db := pgRepo.postgres.DB

	err := db.DropTableIfExists(&model.Application{}, &model.User{}).Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.User{}, &model.Application{}).Error
	if err != nil {
		return err
	}

	err = db.Model(&model.Application{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error
	if err != nil {
		return err
	}

	return nil
}

// User seeders
func seedOneUser() (*model.User, error) {

	db := pgRepo.postgres.DB

	var user = model.User{
		Email:     "johndoe@gmail.com",
		Password:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
	}

	err := db.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func seedUsers() (*[]model.User, error) {

	db := pgRepo.postgres.DB

	var users = []model.User{
		{
			Email:     "johndoe@gmail.com",
			Password:  "johndoe",
			FirstName: "John",
			LastName:  "Doe",
		},
		{
			Email:     "mariodraghi@gmail.com",
			Password:  "mariodraghi",
			FirstName: "Mario",
			LastName:  "Draghi",
		},
	}

	for i := range users {
		err := db.Model(&model.User{}).Create(&users[i]).Error
		if err != nil {
			return nil, err
		}
	}

	return &users, nil
}

// Application seeders
func seedOneApplication() (*model.Application, error) {

	db := pgRepo.postgres.DB

	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	var application = model.Application{
		JobTitle: "Software Engineer Intern",
		Company:  "GoCardless",
		UserID:   user.ID,
	}

	err = db.Model(&model.Application{}).Create(&application).Error
	if err != nil {
		return nil, err
	}

	return &application, nil
}

func seedApplications() (*[]model.Application, error) {

	db := pgRepo.postgres.DB

	users, err := seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	var applications = []model.Application{
		{
			JobTitle: "Software Engineer Intern",
			Company:  "GoCardless",
		},
		{
			JobTitle: "Software Engineer Intern",
			Company:  "Skyscanner",
		},
	}

	for i := range applications {
		applications[i].UserID = (*users)[i].ID
		err = db.Model(&model.Application{}).Create(&applications[i]).Error
		if err != nil {
			return nil, err
		}
	}

	return &applications, nil
}
