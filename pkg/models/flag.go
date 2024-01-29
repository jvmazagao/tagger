package models

import (
	"gorm.io/gorm"
)

type FlagData struct {
	Name        string `json:"name"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
}

type Flag struct {
	gorm.Model
	ID          int
	Name        string
	Enabled     bool
	Description *string
}

type FlagService interface {
	Create(f *FlagData) (Flag, error)
	Delete(id int) error
	Get(id int) (Flag, error)
	GetAll() ([]Flag, error)
	Update(f *FlagData) (Flag, error)
}

type FlagRepository interface {
	InsertFlag(f *Flag) (Flag, error)
	GetFlags() ([]Flag, error)
	DeleteFlag(Id int) error
	GetFlag(Id int) (Flag, error)
	UpdateFlag(f *Flag) (Flag, error)
}
