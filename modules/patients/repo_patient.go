package patients

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aldiramdan/hospital/databases/db/models"
)

type patient_repo struct {
	queryGetAllPatient    string
	queryGetByIdPatient   string
	queryGetByNamePatient string
	queryAddPatient       string
	queryUpdatePatient    string
	queryDeletePatient    string
	db                    *sql.DB
}

func NewRepo(db *sql.DB) *patient_repo {
	return &patient_repo{
		queryGetAllPatient:    "SELECT * FROM patient",
		queryGetByIdPatient:   "SELECT * FROM patient WHERE id=?",
		queryGetByNamePatient: "SELECT * FROM patient WHERE name=?",
		queryAddPatient:       "INSERT INTO patient(id, name, age, address) VALUES(?,?,?,?)",
		queryUpdatePatient:    "UPDATE patient SET name=?, age=?, address=? WHERE id=?",
		queryDeletePatient:    "DELETE FROM patient WHERE id=?",
		db:                    db,
	}
}

func (r *patient_repo) GetAll() (*models.Patients, error) {
	var patient models.Patient
	var patients models.Patients

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, r.queryGetAllPatient)
	if err != nil {
		return nil, errors.New("failed get data patient")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&patient.PatientId, &patient.Name, &patient.Age, &patient.Address, &patient.CreatedAt, &patient.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data patient")
		}
		patients = append(patients, patient)
	}

	return &patients, nil
}

func (r *patient_repo) GetById(id string) (*models.Patient, error) {
	var patient models.Patient

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows := r.db.QueryRowContext(ctx, r.queryGetByIdPatient, id)
	err := rows.Scan(&patient.PatientId, &patient.Name, &patient.Age, &patient.Address, &patient.CreatedAt, &patient.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("data patient not found")
	} else if err != nil {
		return nil, errors.New("failed to scan data patient")
	}

	return &patient, nil
}

func (r *patient_repo) GetByName(name string) (*models.Patient, error) {
	var patient models.Patient

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows := r.db.QueryRowContext(ctx, r.queryGetByNamePatient, name)
	err := rows.Scan(&patient.PatientId, &patient.Name, &patient.Age, &patient.Address, &patient.CreatedAt, &patient.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("data patient not found")
	} else if err != nil {
		return nil, errors.New("failed to scan data patient")
	}

	return &patient, nil
}

func (r *patient_repo) Add(data models.Patient) (*models.Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, r.queryAddPatient, data.PatientId, data.Name, data.Age, data.Address)
	if err != nil {
		return nil, errors.New("failed to add data patient")
	}

	result, err := r.GetById(data.PatientId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *patient_repo) Update(id string, data models.Patient) (*models.Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, r.queryUpdatePatient, data.Name, data.Age, data.Address, data.PatientId)
	if err != nil {
		return nil, errors.New("failed to update data patient")
	}

	result, err := r.GetById(data.PatientId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *patient_repo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.QueryContext(ctx, r.queryDeletePatient, id)
	if err != nil {
		return errors.New("failed get data patient")
	}

	return nil
}
