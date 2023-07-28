package doctors

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
	"github.com/aldiramdan/hospital/interfaces"
)

type doctor_srvc struct {
	repo interfaces.DoctorRepo
}

func NewSrvc(repo interfaces.DoctorRepo) *doctor_srvc {
	return &doctor_srvc{repo}
}

func (s *doctor_srvc) GetAll() *helpers.Response {
	result, err := s.repo.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *doctor_srvc) GetById(id string) *helpers.Response {
	result, err := s.repo.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *doctor_srvc) GetByName(name string) *helpers.Response {
	result, err := s.repo.GetByName(name)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *doctor_srvc) Add(data models.Doctor) *helpers.Response {
	data.DoctorId = helpers.GenerateId()

	result, err := s.repo.Add(data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, true)
}

func (s *doctor_srvc) Update(id string, data models.Doctor) *helpers.Response {
	_, err := s.repo.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	data.DoctorId = id

	result, err := s.repo.Update(id, data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *doctor_srvc) Delete(id string) *helpers.Response {
	_, err := s.repo.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	err = s.repo.Delete(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"message": "data doctor successfully deleted",
	}
	return helpers.GetResponse(response, 200, false)
}

func (s *doctor_srvc) ConvertCSV() *helpers.Response {
	doctors, err := s.repo.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	writer, err := helpers.CreateCSV(os.Getenv("PathDoctorCSV"))
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}
	defer writer.Flush()

	header := []string{"Doctor ID", "Name", "Specialization", "Consultation Fee"}
	writer.Write(header)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, v := range *doctors {
		wg.Add(1)
		go func(v models.Doctor) {
			defer wg.Done()

			id := v.DoctorId
			name := v.Name
			specialization := v.Specialization
			consultationFee := fmt.Sprintf("%.2f", v.ConsultationFee)

			row := []string{id, name, specialization, consultationFee}

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

func (s *doctor_srvc) Download() *helpers.Response {
	file, err := os.Open(os.Getenv("PathDoctorCSV"))
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
