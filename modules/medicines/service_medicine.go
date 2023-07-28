package medicines

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
	"github.com/aldiramdan/hospital/interfaces"
)

type medicine_srvc struct {
	repo interfaces.MedicineRepo
}

func NewSrvc(repo interfaces.MedicineRepo) *medicine_srvc {
	return &medicine_srvc{repo}
}

func (s *medicine_srvc) GetAll() *helpers.Response {
	result, err := s.repo.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *medicine_srvc) GetById(id string) *helpers.Response {
	result, err := s.repo.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *medicine_srvc) GetByName(name string) *helpers.Response {
	result, err := s.repo.GetByName(name)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *medicine_srvc) Add(data models.Medicine) *helpers.Response {
	data.MedicineId = helpers.GenerateId()

	result, err := s.repo.Add(data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *medicine_srvc) Update(id string, data models.Medicine) *helpers.Response {
	_, err := s.repo.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	data.MedicineId = id

	result, err := s.repo.Update(id, data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *medicine_srvc) Delete(id string) *helpers.Response {
	_, err := s.repo.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	err = s.repo.Delete(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"message": "data medicine successfully deleted",
	}
	return helpers.GetResponse(response, 200, false)
}

func (s *medicine_srvc) ConvertCSV() *helpers.Response {
	medicines, err := s.repo.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	writer, err := helpers.CreateCSV(os.Getenv("PathMedicineCSV"))
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}
	defer writer.Flush()

	header := []string{"Medicine ID", "Name", "Price"}
	writer.Write(header)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, v := range *medicines {
		wg.Add(1)
		go func(v models.Medicine) {
			defer wg.Done()

			id := v.MedicineId
			name := v.Name
			price := fmt.Sprintf("%.2f", v.Price)

			row := []string{id, name, price}

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

func (s *medicine_srvc) Download() *helpers.Response {
	file, err := os.Open(os.Getenv("PathMedicineCSV"))
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
