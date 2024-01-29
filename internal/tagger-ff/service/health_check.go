package service

import (
	"tagger-ff/pkg/models"
)

func IsAlive() ([]models.HealthCheck, error) {
	serviceStatus := make([]models.HealthCheck, 0)
	serviceStatus = append(serviceStatus, models.HealthCheck{
		Service: "tagger-ff",
		Status:  true,
	})
	return serviceStatus, nil
}
