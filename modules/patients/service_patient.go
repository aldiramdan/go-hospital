package patients

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
	"github.com/aldiramdan/hospital/interfaces"
)

type patient_srvc struct {
	repo interfaces.PatientRepo
}

func NewSrvc(repo interfaces.PatientRepo) *patient_srvc {
	return &patient_srvc{repo}
}

func (s *patient_srvc) GetAll() *helpers.Response {
	result, err := s.repo.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *patient_srvc) GetById(id string) *helpers.Response {
	result, err := s.repo.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *patient_srvc) GetByName(name string) *helpers.Response {
	result, err := s.repo.GetByName(name)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *patient_srvc) Add(data models.Patient) *helpers.Response {
	data.PatientId = helpers.GenerateId()

	result, err := s.repo.Add(data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *patient_srvc) Update(id string, data models.Patient) *helpers.Response {
	_, err := s.repo.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	data.PatientId = id

	result, err := s.repo.Update(id, data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *patient_srvc) Delete(id string) *helpers.Response {
	_, err := s.repo.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	err = s.repo.Delete(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"message": "data patient successfully deleted",
	}
	return helpers.GetResponse(response, 200, false)
}

func (s *patient_srvc) ConvertCSV() *helpers.Response {
	patients, err := s.repo.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	writer, err := helpers.CreateCSV(os.Getenv("PathPatientCSV"))
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}
	defer writer.Flush()

	header := []string{"Patient ID", "Name", "Age", "Address"}
	writer.Write(header)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, v := range *patients {
		wg.Add(1)
		go func(v models.Patient) {
			defer wg.Done()

			id := v.PatientId
			name := v.Name
			age := fmt.Sprintf("%d", v.Age)
			address := v.Address

			row := []string{id, name, age, address}

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

func (s *patient_srvc) Download() *helpers.Response {
	file, err := os.Open(os.Getenv("PathPatientCSV"))
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
