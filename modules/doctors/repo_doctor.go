package doctors

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aldiramdan/hospital/databases/db/models"
)

type doctor_repo struct {
	queryGetAllDoctor    string
	queryGetByIdDoctor   string
	queryGetByNameDoctor string
	queryAddDoctor       string
	queryUpdateDoctor    string
	queryDeleteDoctor    string
	db                   *sql.DB
}

func NewRepo(db *sql.DB) *doctor_repo {
	return &doctor_repo{
		queryGetAllDoctor:    "SELECT * FROM doctor",
		queryGetByIdDoctor:   "SELECT * FROM doctor WHERE id=?",
		queryGetByNameDoctor: "SELECT * FROM doctor WHERE name=?",
		queryAddDoctor:       "INSERT INTO doctor(id, name, specialization, consultation_fee) VALUES(?,?,?,?)",
		queryUpdateDoctor:    "UPDATE doctor SET name=?, specialization=?, consultation_fee=? WHERE id=? ",
		queryDeleteDoctor:    "DELETE FROM doctor WHERE id=?",
		db:                   db,
	}
}

func (r *doctor_repo) GetAll() (*models.Doctors, error) {
	var doctor models.Doctor
	var doctors models.Doctors

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, r.queryGetAllDoctor)
	if err != nil {
		return nil, errors.New("failed get data doctor")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&doctor.DoctorId, &doctor.Name, &doctor.Specialization, &doctor.ConsultationFee, &doctor.CreatedAt, &doctor.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data doctor")
		}
		doctors = append(doctors, doctor)
	}

	return &doctors, nil
}

func (r *doctor_repo) GetById(id string) (*models.Doctor, error) {
	var doctor models.Doctor

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows := r.db.QueryRowContext(ctx, r.queryGetByIdDoctor, id)
	err := rows.Scan(&doctor.DoctorId, &doctor.Name, &doctor.Specialization, &doctor.ConsultationFee, &doctor.CreatedAt, &doctor.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("data doctor not found")
	} else if err != nil {
		return nil, errors.New("failed to scan data doctor")
	}

	return &doctor, nil
}

func (r *doctor_repo) GetByName(name string) (*models.Doctor, error) {
	var doctor models.Doctor

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows := r.db.QueryRowContext(ctx, r.queryGetByNameDoctor, name)
	err := rows.Scan(&doctor.DoctorId, &doctor.Name, &doctor.Specialization, &doctor.ConsultationFee, &doctor.CreatedAt, &doctor.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("data doctor not found")
	} else if err != nil {
		return nil, errors.New("failed to scan data doctor")
	}

	return &doctor, nil
}

func (r *doctor_repo) Add(data models.Doctor) (*models.Doctor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, r.queryAddDoctor, &data.DoctorId, &data.Name, &data.Specialization, &data.ConsultationFee)
	if err != nil {
		return nil, errors.New("failed to add data doctor")
	}

	result, err := r.GetById(data.DoctorId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *doctor_repo) Update(id string, data models.Doctor) (*models.Doctor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, r.queryUpdateDoctor, &data.Name, &data.Specialization, &data.ConsultationFee, &data.DoctorId)
	if err != nil {
		return nil, errors.New("failed to update data doctor")
	}

	result, err := r.GetById(data.DoctorId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *doctor_repo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.QueryContext(ctx, r.queryDeleteDoctor, id)
	if err != nil {
		return errors.New("failed get data doctor")
	}

	return nil
}
