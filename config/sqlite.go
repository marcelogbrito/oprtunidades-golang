package config

import (
	"os"

	"github.com/marcelogbrito/oprtunidades-golang/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"
	logger := GetLogger("sqlite")
	//check if database exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// create database, file and directory
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
	//Create
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqllite opening error: %v", err)
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqllite automigration error: %v", err)
		return nil, err
	}
	//return db
	return db, nil
}
