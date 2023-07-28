package treatments

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
	"github.com/aldiramdan/hospital/interfaces"
)

type treatment_srvc struct {
	repo_patient   interfaces.PatientRepo
	repo_disease   interfaces.DiseaseRepo
	repo_doctor    interfaces.DoctorRepo
	repo_treatment interfaces.TreatmentRepo
}

func NewSrvc(repo_patient interfaces.PatientRepo, repo_disease interfaces.DiseaseRepo, repo_doctor interfaces.DoctorRepo, repo_treatment interfaces.TreatmentRepo) *treatment_srvc {
	return &treatment_srvc{
		repo_patient:   repo_patient,
		repo_disease:   repo_disease,
		repo_doctor:    repo_doctor,
		repo_treatment: repo_treatment,
	}
}

func (s *treatment_srvc) GetAll() *helpers.Response {
	result, err := s.repo_treatment.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *treatment_srvc) GetById(id string) *helpers.Response {
	result, err := s.repo_treatment.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *treatment_srvc) GetByNamePatient(name string) *helpers.Response {
	result, err := s.repo_treatment.GetByNamePatient(name)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, false)
}

func (s *treatment_srvc) Add(data models.Treatment) *helpers.Response {
	patient, err := s.repo_patient.GetById(data.Patient.PatientId)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	doctor, err := s.repo_doctor.GetById(data.Doctor.DoctorId)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	dupeName, err := s.repo_disease.DuplicateName(data.Disease)
	if dupeName {
		if err != nil {
			return helpers.GetResponse(err.Error(), 500, true)
		}
	}

	for i := range data.Disease {
		disease, err := s.repo_disease.GetByName(data.Disease[i].Name)
		if err != nil {
			return helpers.GetResponse(err.Error(), 500, true)
		}
		data.Disease[i] = *disease
	}

	fee := helpers.ServiceFeeDoctor(data.Disease, doctor.ConsultationFee)

	data = models.Treatment{
		TreatmentId: helpers.GenerateId(),
		Patient:     *patient,
		Disease:     data.Disease,
		Doctor:      *doctor,
		ServiceFee:  fee,
	}

	result, err := s.repo_treatment.Add(data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, true)
}

func (s *treatment_srvc) Update(id string, data models.Treatment) *helpers.Response {
	_, err := s.repo_treatment.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	patient, err := s.repo_patient.GetById(data.Patient.PatientId)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	doctor, err := s.repo_doctor.GetById(data.Doctor.DoctorId)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	dupeName, err := s.repo_disease.DuplicateName(data.Disease)
	if dupeName {
		if err != nil {
			return helpers.GetResponse(err.Error(), 500, true)
		}
	}

	for i := range data.Disease {
		disease, err := s.repo_disease.GetByName(data.Disease[i].Name)
		if err != nil {
			return helpers.GetResponse(err.Error(), 500, true)
		}
		data.Disease[i] = *disease
	}

	fee := helpers.ServiceFeeDoctor(data.Disease, doctor.ConsultationFee)

	data = models.Treatment{
		TreatmentId: helpers.GenerateId(),
		Patient:     *patient,
		Disease:     data.Disease,
		Doctor:      *doctor,
		ServiceFee:  fee,
	}

	datatreatment, err := s.repo_treatment.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	datatreatment.TreatmentId = id
	data.TreatmentId = id
	datatreatment = &data

	result, err := s.repo_treatment.Update(id, data)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	return helpers.GetResponse(result, 200, true)
}

func (s *treatment_srvc) Delete(id string) *helpers.Response {
	_, err := s.repo_treatment.GetById(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	err = s.repo_treatment.Delete(id)
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"message": "data treatment successfully deleted",
	}
	return helpers.GetResponse(response, 200, false)
}

func (s *treatment_srvc) ConvertCSV() *helpers.Response {
	treatments, err := s.repo_treatment.GetAll()
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}

	writer, err := helpers.CreateCSV(os.Getenv("PathtreatmentCSV"))
	if err != nil {
		return helpers.GetResponse(err.Error(), 500, true)
	}
	defer writer.Flush()

	header := []string{"treatment ID", "Patient Name", "Patient Age", "Patient Address", "Disease Name", "Medicine Name", "Doctor Name", "Service Fee"}
	writer.Write(header)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, v := range *treatments {
		wg.Add(1)
		go func(v models.Treatment) {
			defer wg.Done()

			treatmentId := v.TreatmentId
			patientName := v.Patient.Name
			patientAge := fmt.Sprintf("%d", v.Patient.Age)
			patientAddress := v.Patient.Address
			doctorName := v.Doctor.Name
			serviceFee := fmt.Sprintf("%.2f", v.ServiceFee)

			diseaseSlice := make([]string, len(v.Disease))
			medicineSlice := make([]string, len(v.Disease))

			for i, p := range v.Disease {
				diseaseSlice[i] = p.Name
				medicineSlice[i] = p.Medicine.Name
			}

			resultDisease := strings.Join(diseaseSlice, ", ")
			resultMedicine := strings.Join(medicineSlice, ", ")

			row := []string{treatmentId, patientName, patientAge, patientAddress, resultDisease, resultMedicine, doctorName, serviceFee}
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

func (s *treatment_srvc) Download() *helpers.Response {
	file, err := os.Open(os.Getenv("PathtreatmentCSV"))
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
