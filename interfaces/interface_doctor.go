package interfaces

import (
	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
)

type DoctorRepo interface {
	GetAll() (*models.Doctors, error)
	GetById(id string) (*models.Doctor, error)
	GetByName(name string) (*models.Doctor, error)
	Add(data models.Doctor) (*models.Doctor, error)
	Update(id string, data models.Doctor) (*models.Doctor, error)
	Delete(id string) error
}

type DoctorSrvc interface {
	GetAll() *helpers.Response
	GetById(id string) *helpers.Response
	GetByName(name string) *helpers.Response
	Add(data models.Doctor) *helpers.Response
	Update(id string, data models.Doctor) *helpers.Response
	Delete(id string) *helpers.Response
	ConvertCSV() *helpers.Response
	Download() *helpers.Response
}
