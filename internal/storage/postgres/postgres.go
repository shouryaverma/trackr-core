package postgres

import (
	"fmt"
	"log"

	"github.com/amaraliou/trackr-v2/internal/model"
	"github.com/amaraliou/trackr-v2/pkg/logger"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connection ...
type Connection struct {
	DB     *gorm.DB
	logger logger.Logger
}

// Repository ...
type Repository struct {
	postgres *Connection
}

// NewRepository ...
func NewRepository(connection *Connection) (*Repository, error) {
	return &Repository{
		postgres: connection,
	}, nil
}

// NewConnection ...
func NewConnection(config *Config, logger logger.Logger) (*Connection, error) {
	username := config.User
	password := config.Password
	dbName := config.DbName
	dbHost := config.Host

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&model.User{},
	)

	connection := Connection{
		DB:     db,
		logger: logger,
	}

	return &connection, nil
}

// Close ...
func (c *Connection) Close() {
	err := c.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
