package interfaces

import (
	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
)

type MedicineRepo interface {
	GetAll() (*models.Medicines, error)
	GetById(id string) (*models.Medicine, error)
	GetByName(name string) (*models.Medicine, error)
	Add(data models.Medicine) (*models.Medicine, error)
	Update(id string, data models.Medicine) (*models.Medicine, error)
	Delete(id string) error
}

type MedicineSrvc interface {
	GetAll() *helpers.Response
	GetById(id string) *helpers.Response
	GetByName(name string) *helpers.Response
	Add(data models.Medicine) *helpers.Response
	Update(id string, data models.Medicine) *helpers.Response
	Delete(id string) *helpers.Response
	ConvertCSV() *helpers.Response
	Download() *helpers.Response
}
