package interfaces

import (
	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
)

type PatientRepo interface {
	GetAll() (*models.Patients, error)
	GetById(id string) (*models.Patient, error)
	GetByName(name string) (*models.Patient, error)
	Add(data models.Patient) (*models.Patient, error)
	Update(id string, data models.Patient) (*models.Patient, error)
	Delete(id string) error
}

type PatientSrvc interface {
	GetAll() *helpers.Response
	GetById(id string) *helpers.Response
	GetByName(name string) *helpers.Response
	Add(data models.Patient) *helpers.Response
	Update(id string, data models.Patient) *helpers.Response
	Delete(id string) *helpers.Response
	ConvertCSV() *helpers.Response
	Download() *helpers.Response
}
