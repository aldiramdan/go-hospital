package interfaces

import (
	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
)

type TreatmentRepo interface {
	GetAll() (*models.Treatments, error)
	GetById(id string) (*models.Treatment, error)
	GetByNamePatient(name string) (*models.Treatments, error)
	Add(data models.Treatment) (*models.Treatment, error)
	Update(id string, data models.Treatment) (*models.Treatment, error)
	Delete(id string) error
}

type TreatmentSrvc interface {
	GetAll() *helpers.Response
	GetById(id string) *helpers.Response
	GetByNamePatient(name string) *helpers.Response
	Add(data models.Treatment) *helpers.Response
	Update(id string, data models.Treatment) *helpers.Response
	Delete(id string) *helpers.Response
	ConvertCSV() *helpers.Response
	Download() *helpers.Response
}
