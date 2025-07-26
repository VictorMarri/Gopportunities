package config

import (
	"github.com/VictorMarri/Gopportunities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializeSQLite() (*gorm.DB, error) /*What will return on this method*/ {
	logger := GetLogger("Initializing SQLite instance")
	dbPath := "./db/main.db"

	//Check if database already exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found. Creating...")
		//Create database file and directory
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

	//Creating database and connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error opening SQLite instance: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{}) //Passing the struct opening pointer to create database migration based on it

	if err != nil {
		logger.Errorf("Error migrating database: %v", err)
		return nil, err
	}

	return db, nil
}
