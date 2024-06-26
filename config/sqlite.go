package config

import (
	"os"

	"github.com/gustasousagh/jobs-api-golang/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating")
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
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrorF("sqlite error Opening: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schema.Opening{}, &schema.User{})
	if err != nil {
		logger.ErrorF("sqlite Automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
