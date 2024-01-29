package database

import (
	"tagger-ff/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=example dbname=recordings port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if database != nil {
		return database, nil
	}

	database = db

	return database, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Flag{})
}

func GetDatabase() *gorm.DB {
	if database == nil {
		database, _ = Connect()
	}

	return database
}
