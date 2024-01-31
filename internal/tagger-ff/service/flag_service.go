package service

import (
	"tagger-ff/internal/tagger-ff/repo/postgres"
	"tagger-ff/pkg/models"
)

func GetAllFlags() ([]models.Flag, error) {
	return postgres.GetFlags()
}

func CreateFlag(f *models.FlagData) (models.Flag, error) {
	flag := models.Flag{
		Description: f.Description,
		Name:        f.Name,
		Enabled:     f.Enabled,
	}

	return postgres.InsertFlag(flag)
}
