package models

type HealthCheck struct {
	Service string
	Status  bool
}

type HealthCheckService interface {
	IsAlive() ([]HealthCheck, error)
}
