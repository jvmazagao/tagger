package postgres

import (
	"tagger-ff/internal/database"
	"tagger-ff/pkg/models"
)

func GetFlags() ([]models.Flag, error) {
	db := database.GetDatabase()
	var flags []models.Flag
	result := db.Find(&flags)
	return flags, result.Error
}

func InsertFlag(f models.Flag) (models.Flag, error) {
	db := database.GetDatabase()
	res := db.Create(f)
	return f, res.Error
}
