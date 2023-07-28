package interfaces

import (
	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
)

type DiseaseRepo interface {
	GetAll() (*models.Diseases, error)
	GetById(id string) (*models.Disease, error)
	GetByName(name string) (*models.Disease, error)
	Add(data models.Disease) (*models.Disease, error)
	Update(id string, data models.Disease) (*models.Disease, error)
	Delete(id string) error
	DuplicateName(data models.Diseases) (bool, error)
}

type DiseaseSrvc interface {
	GetAll() *helpers.Response
	GetById(id string) *helpers.Response
	GetByName(name string) *helpers.Response
	Add(data models.Disease) *helpers.Response
	Update(id string, data models.Disease) *helpers.Response
	Delete(id string) *helpers.Response
	ConvertCSV() *helpers.Response
	Download() *helpers.Response
}
