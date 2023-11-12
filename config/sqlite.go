package config

import (
	"os"

	"github.com/maikondouglas/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the database is already initialized
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")

		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error creating database: %v", err)
		return nil, err
	}

	// Migreate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error auto-migrate: %v", err)
		return nil, err
	}

	// Create database
	return db, nil
}
