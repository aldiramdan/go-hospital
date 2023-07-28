package treatments

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aldiramdan/hospital/databases/db/models"
)

type treatment_repo struct {
	queryGetAllTreatment         string
	queryGetByIdTreatment        string
	queryGetByNamePatient        string
	queryAddTreatment            string
	queryUpdateTreatment         string
	queryDeleteTreatment         string
	queryGetAllTreatmentDisease  string
	queryGetByIdTreatmentDisease string
	queryAddTreatmentDisease     string
	queryDeleteTreatmentDisease  string
	db                           *sql.DB
}

func NewRepo(db *sql.DB) *treatment_repo {
	return &treatment_repo{
		queryGetAllTreatment: `SELECT t.id, t.service_fee,
									pt.id, pt.name, pt.age, pt.address, pt.created_at, pt.updated_at,
									dc.id, dc.name, dc.specialization, dc.consultation_fee, dc.created_at, dc.updated_at,
									t.created_at, t.updated_at
								FROM treatment AS t
								JOIN patient AS pt ON t.patient_id = pt.id
								JOIN doctor AS dc ON t.doctor_id = dc.id`,
		queryGetAllTreatmentDisease: `SELECT hd.treatment_id, hd.disease_id,
									d.id, d.name, d.created_at, d.updated_at,
									m.id, m.name, m.price, m.created_at, m.updated_at
								FROM treatment_disease AS hd
								JOIN treatment t ON hd.treatment_id = t.id
								JOIN disease d ON hd.disease_id = d.id
								JOIN medicine m ON d.medicine_id = m.id`,
		queryGetByIdTreatment: `SELECT t.id, t.service_fee,
									pt.id, pt.name, pt.age, pt.address, pt.created_at, pt.updated_at,
									dc.id, dc.name, dc.specialization, dc.consultation_fee, dc.created_at, dc.updated_at,
									t.created_at, t.updated_at
								FROM treatment AS t
								JOIN patient AS pt ON t.patient_id = pt.id
								JOIN doctor AS dc ON t.doctor_id = dc.id
								WHERE t.id=?`,
		queryGetByIdTreatmentDisease: `SELECT hd.treatment_id, hd.disease_id,
											d.id, d.name, d.created_at, d.updated_at,
											m.id, m.name, m.price, m.created_at, m.updated_at
										FROM treatment_disease AS hd
										JOIN treatment t ON hd.treatment_id = t.id
										JOIN disease d ON hd.disease_id = d.id
										JOIN medicine m ON d.medicine_id = m.id
										WHERE t.id=?`,
		queryGetByNamePatient: `SELECT t.id, t.service_fee,
									pt.id, pt.name, pt.age, pt.address, pt.created_at, pt.updated_at,
									dc.id, dc.name, dc.specialization, dc.consultation_fee, dc.created_at, dc.updated_at,
									t.created_at, t.updated_at
								FROM treatment AS t
								JOIN patient AS pt ON t.patient_id = pt.id
								JOIN doctor AS dc ON t.doctor_id = dc.id
								WHERE pt.name LIKE ?`,
		queryAddTreatment:           "INSERT INTO treatment(id, patient_id, doctor_id, service_fee) VALUES(?,?,?,?)",
		queryUpdateTreatment:        "UPDATE treatment SET patient_id=?, doctor_id=?, service_fee=? WHERE id=?",
		queryDeleteTreatment:        "DELETE FROM treatment WHERE id=?",
		queryAddTreatmentDisease:    "INSERT INTO treatment_disease(treatment_id, disease_id) VALUES(?,?)",
		queryDeleteTreatmentDisease: "DELETE FROM treatment_disease WHERE treatment_id=?",
		db:                          db,
	}
}

func (r *treatment_repo) GetAll() (*models.Treatments, error) {
	var treatment models.Treatment
	var treatments models.Treatments
	var disease models.Disease

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rowsTreatment, err := r.db.QueryContext(ctx, r.queryGetAllTreatment)
	if err != nil {
		return nil, errors.New("failed get data treatment")
	}
	defer rowsTreatment.Close()

	for rowsTreatment.Next() {
		err := rowsTreatment.Scan(
			&treatment.TreatmentId, &treatment.ServiceFee,
			&treatment.Patient.PatientId, &treatment.Patient.Name, &treatment.Patient.Age, &treatment.Patient.Address, &treatment.Patient.CreatedAt, &treatment.Patient.UpdatedAt,
			&treatment.Doctor.DoctorId, &treatment.Doctor.Name, &treatment.Doctor.Specialization, &treatment.Doctor.ConsultationFee, &treatment.Doctor.CreatedAt, &treatment.Doctor.UpdatedAt,
			&treatment.CreatedAt, &treatment.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data treatment")
		}

		treatments = append(treatments, treatment)
	}

	rowsDisease, err := r.db.QueryContext(ctx, r.queryGetAllTreatmentDisease)
	if err != nil {
		return nil, errors.New("failed get data treatment and disease")

	}
	defer rowsDisease.Close()

	for rowsDisease.Next() {
		err := rowsDisease.Scan(
			&treatment.TreatmentId,
			&disease.DiseaseId, &disease.DiseaseId, &disease.Name, &disease.CreatedAt, &disease.UpdatedAt,
			&disease.Medicine.MedicineId, &disease.Medicine.Name, &disease.Medicine.Price, &disease.Medicine.CreatedAt, &disease.Medicine.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data treatment and disease")
		}

		for i, v := range treatments {
			if v.TreatmentId == treatment.TreatmentId {
				treatments[i].Disease = append(treatments[i].Disease, disease)
				break
			}
		}
	}

	return &treatments, nil
}

func (r *treatment_repo) GetById(id string) (*models.Treatment, error) {
	var treatment models.Treatment
	var disease models.Disease

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rowsTreatment := r.db.QueryRowContext(ctx, r.queryGetByIdTreatment, id)
	err := rowsTreatment.Scan(
		&treatment.TreatmentId, &treatment.ServiceFee,
		&treatment.Patient.PatientId, &treatment.Patient.Name, &treatment.Patient.Age, &treatment.Patient.Address, &treatment.Patient.CreatedAt, &treatment.Patient.UpdatedAt,
		&treatment.Doctor.DoctorId, &treatment.Doctor.Name, &treatment.Doctor.Specialization, &treatment.Doctor.ConsultationFee, &treatment.Doctor.CreatedAt, &treatment.Doctor.UpdatedAt,
		&treatment.CreatedAt, &treatment.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("data treatment not found")
	} else if err != nil {
		return nil, errors.New("failed to scan data treatment")
	}

	treatment.Disease = models.Diseases{}

	rowsDisease, err := r.db.QueryContext(ctx, r.queryGetByIdTreatmentDisease, id)
	if err != nil {
		return nil, errors.New("failed get data treatment and disease")
	}
	defer rowsDisease.Close()

	for rowsDisease.Next() {
		err := rowsDisease.Scan(
			&treatment.TreatmentId,
			&disease.DiseaseId, &disease.DiseaseId, &disease.Name, &disease.CreatedAt, &disease.UpdatedAt,
			&disease.Medicine.MedicineId, &disease.Medicine.Name, &disease.Medicine.Price, &disease.Medicine.CreatedAt, &disease.Medicine.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data treatment and disease")
		}

		treatment.Disease = append(treatment.Disease, disease)
	}

	return &treatment, nil
}

func (r *treatment_repo) GetByNamePatient(name string) (*models.Treatments, error) {
	var treatment models.Treatment
	var treatments models.Treatments
	var disease models.Disease

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rowsTreatment, err := r.db.QueryContext(ctx, r.queryGetByNamePatient, name)
	if err != nil {
		return nil, errors.New("failed get data treatment")
	}
	defer rowsTreatment.Close()

	for rowsTreatment.Next() {
		err := rowsTreatment.Scan(
			&treatment.TreatmentId, &treatment.ServiceFee,
			&treatment.Patient.PatientId, &treatment.Patient.Name, &treatment.Patient.Age, &treatment.Patient.Address, &treatment.Patient.CreatedAt, &treatment.Patient.UpdatedAt,
			&treatment.Doctor.DoctorId, &treatment.Doctor.Name, &treatment.Doctor.Specialization, &treatment.Doctor.ConsultationFee, &treatment.Doctor.CreatedAt, &treatment.Doctor.UpdatedAt,
			&treatment.CreatedAt, &treatment.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data treatment")
		}

		treatments = append(treatments, treatment)
	}

	rowsDisease, err := r.db.QueryContext(ctx, r.queryGetAllTreatmentDisease)
	if err != nil {
		return nil, errors.New("failed get data treatment and disease")
	}
	defer rowsDisease.Close()

	for rowsDisease.Next() {
		err := rowsDisease.Scan(
			&treatment.TreatmentId,
			&disease.DiseaseId, &disease.DiseaseId, &disease.Name, &disease.CreatedAt, &disease.UpdatedAt,
			&disease.Medicine.MedicineId, &disease.Medicine.Name, &disease.Medicine.Price, &disease.Medicine.CreatedAt, &disease.Medicine.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data treatment and disease")
		}

		for i, v := range treatments {
			if v.TreatmentId == treatment.TreatmentId {
				treatments[i].Disease = append(treatments[i].Disease, disease)
				break
			}
		}
	}

	return &treatments, nil
}

func (r *treatment_repo) Add(data models.Treatment) (*models.Treatment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, r.queryAddTreatment, data.TreatmentId, data.Patient.PatientId, data.Doctor.DoctorId, data.ServiceFee)
	if err != nil {
		return nil, errors.New("failed to add data treatment")
	}

	stmt, err := tx.PrepareContext(ctx, r.queryAddTreatmentDisease)
	if err != nil {
		return nil, errors.New("failed to prepare statement for adding diseases")
	}
	defer stmt.Close()

	for _, v := range data.Disease {
		_, err = stmt.ExecContext(ctx, data.TreatmentId, v.DiseaseId)
		if err != nil {
			return nil, errors.New("failed to add data disease")
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	result, err := r.GetById(data.TreatmentId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *treatment_repo) Update(id string, data models.Treatment) (*models.Treatment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, r.queryUpdateTreatment, data.Patient.PatientId, data.Doctor.DoctorId, data.ServiceFee, data.TreatmentId)
	if err != nil {
		return nil, errors.New("failed to update data treatment")
	}

	_, err = tx.ExecContext(ctx, r.queryDeleteTreatmentDisease, id)
	if err != nil {
		return nil, errors.New("failed to prepare statement for delete diseases")
	}

	stmt, err := tx.PrepareContext(ctx, r.queryAddTreatmentDisease)
	if err != nil {
		return nil, errors.New("failed to prepare statement for adding diseases")
	}
	defer stmt.Close()

	for _, v := range data.Disease {
		_, err = stmt.ExecContext(ctx, data.TreatmentId, v.DiseaseId)
		if err != nil {
			return nil, errors.New("failed to add data disease")
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	result, err := r.GetById(data.TreatmentId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *treatment_repo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, r.queryDeleteTreatmentDisease, id)
	if err != nil {
		return errors.New("failed to prepare statement for delete diseases")
	}

	_, err = tx.ExecContext(ctx, r.queryDeleteTreatment, id)
	if err != nil {
		return errors.New("failed to delete data treatment")
	}

	if err := tx.Commit(); err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}
