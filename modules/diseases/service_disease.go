package diseases

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
	"github.com/aldiramdan/hospital/interfaces"
)

type disease_srvc struct {
	repo_disease  interfaces.DiseaseRepo
	repo_medicine interfaces.MedicineRepo
}

func NewSrvc(repo_disease interfaces.DiseaseRepo, repo_medicine interfaces.MedicineRepo) *disease_srvc {
	return &disease_srvc{
		repo_disease:  repo_disease,
		repo_medicine: repo_medicine,
	}
}

func (s *disease_srvc) GetAll() *helpers.Response {
	result, err := s.repo_disease.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *disease_srvc) GetById(id string) *helpers.Response {
	result, err := s.repo_disease.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *disease_srvc) GetByName(name string) *helpers.Response {
	result, err := s.repo_disease.GetByName(name)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *disease_srvc) Add(data models.Disease) *helpers.Response {
	dataMedicine, err := s.repo_medicine.GetByName(data.Medicine.Name)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	data.DiseaseId = helpers.GenerateId()
	data.Medicine = *dataMedicine

	result, err := s.repo_disease.Add(data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *disease_srvc) Update(id string, data models.Disease) *helpers.Response {
	_, err := s.repo_disease.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	dataMedicine, err := s.repo_medicine.GetByName(data.Medicine.Name)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	data.DiseaseId = id
	data.Medicine = *dataMedicine

	result, err := s.repo_disease.Update(id, data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *disease_srvc) Delete(id string) *helpers.Response {
	_, err := s.repo_disease.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	err = s.repo_disease.Delete(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"message": "data disease successfully deleted",
	}
	return helpers.GetResponse(response, 200, false)
}

func (s *disease_srvc) ConvertCSV() *helpers.Response {
	diseases, err := s.repo_disease.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	writer, err := helpers.CreateCSV(os.Getenv("PathDiseaseCSV"))
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}
	defer writer.Flush()

	header := []string{"Disease ID", "Disease Name", "Medicine ID", "Medicine Name", "Medicine Price"}
	writer.Write(header)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, v := range *diseases {
		wg.Add(1)
		go func(v models.Disease) {
			defer wg.Done()

			diseaseId := v.DiseaseId
			diseaseName := v.Name
			medicineId := v.Medicine.MedicineId
			medicineName := v.Medicine.Name
			medicinePrice := fmt.Sprintf("%.2f", v.Medicine.Price)

			row := []string{diseaseId, diseaseName, medicineId, medicineName, medicinePrice}

			mutex.Lock()
			defer mutex.Unlock()

			err := writer.Write(row)
			if err != nil {
				log.Println(err)
			}

		}(v)
	}

	wg.Wait()

	return helpers.GetResponse(map[string]interface{}{"message": "successfully convert JSON to CSV"}, 200, false)
}

func (s *disease_srvc) Download() *helpers.Response {
	file, err := os.Open(os.Getenv("PathDiseaseCSV"))
	if err != nil {
		return helpers.GetResponse(err.Error(), 404, true)

	}
	defer file.Close()

	_, err = file.Stat()
	if err != nil {
		return helpers.GetResponse(err.Error(), 404, true)
	}

	return helpers.GetResponse(map[string]interface{}{"message": "file found, ready for download"}, 200, false)
}
