package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/cybercoder/restbill/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Init initializes the database connection
func Init() *gorm.DB {
	dbOnce.Do(func() {

		cfg := Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		}

		// Construct the connection string
		dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)

		// Connect to the PostgreSQL database
		var err error
		db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
		if err != nil {
			logger.Fatal(err)
		}
	})
	return db
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	if db == nil {
		logger.Panic("database not initialized - call Init() first")
	}
	return db
}

// WithTransaction executes a function within a transaction
func WithTransaction(fn func(tx *gorm.DB) error) error {
	return GetDB().Transaction(fn)
}
