package medicines

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aldiramdan/hospital/databases/db/models"
)

type medicine_repo struct {
	queryGetAllMedicine    string
	queryGetByIdMedicine   string
	queryGetByNameMedicine string
	queryAddMedicine       string
	queryUpdateMedicine    string
	queryDeleteMedicine    string
	queryCountMedicine     string
	db                     *sql.DB
}

func NewRepo(db *sql.DB) *medicine_repo {
	return &medicine_repo{
		queryGetAllMedicine:    "SELECT * FROM medicine",
		queryGetByIdMedicine:   "SELECT * FROM medicine WHERE id=?",
		queryGetByNameMedicine: "SELECT * FROM medicine WHERE name=?",
		queryAddMedicine:       "INSERT INTO medicine(id, name, price) VALUES(?,?,?)",
		queryUpdateMedicine:    "UPDATE medicine SET name=?, price=? WHERE id=?",
		queryDeleteMedicine:    "DELETE FROM medicine WHERE id=?",
		queryCountMedicine:     "SELECT COUNT(*) FROM medicine WHERE name=?",
		db:                     db,
	}
}

func (r *medicine_repo) GetAll() (*models.Medicines, error) {
	var medicine models.Medicine
	var medicines models.Medicines

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, r.queryGetAllMedicine)
	if err != nil {
		return nil, errors.New("failed get data medicine")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&medicine.MedicineId, &medicine.Name, &medicine.Price, &medicine.CreatedAt, &medicine.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data medicine")
		}
		medicines = append(medicines, medicine)
	}

	return &medicines, nil
}

func (r *medicine_repo) GetById(id string) (*models.Medicine, error) {
	var medicine models.Medicine

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows := r.db.QueryRowContext(ctx, r.queryGetByIdMedicine, id)
	err := rows.Scan(&medicine.MedicineId, &medicine.Name, &medicine.Price, &medicine.CreatedAt, &medicine.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("data medicine not found")
	} else if err != nil {
		return nil, errors.New("failed to scan data medicine")
	}

	return &medicine, nil
}

func (r *medicine_repo) GetByName(name string) (*models.Medicine, error) {
	var medicine models.Medicine

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows := r.db.QueryRowContext(ctx, r.queryGetByNameMedicine, name)
	err := rows.Scan(&medicine.MedicineId, &medicine.Name, &medicine.Price, &medicine.CreatedAt, &medicine.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("data medicine not found")
	} else if err != nil {
		return nil, errors.New("failed to scan data medicine")
	}

	return &medicine, nil
}

func (r *medicine_repo) Add(data models.Medicine) (*models.Medicine, error) {
	if exists := r.IsExistsByName(data.Name); exists {
		return nil, errors.New("name medicine has been exists")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, r.queryAddMedicine, data.MedicineId, data.Name, data.Price)
	if err != nil {
		return nil, errors.New("failed to add data medicine")
	}

	result, err := r.GetById(data.MedicineId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *medicine_repo) Update(id string, data models.Medicine) (*models.Medicine, error) {
	if exists := r.IsExistsByName(data.Name); exists {
		return nil, errors.New("name medicine has been exists")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, r.queryUpdateMedicine, data.Name, data.Price, data.MedicineId)
	if err != nil {
		return nil, errors.New("failed to update data medicine")
	}

	result, err := r.GetById(data.MedicineId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *medicine_repo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.QueryContext(ctx, r.queryDeleteMedicine, id)
	if err != nil {
		return errors.New("failed get data medicine")
	}

	return nil
}

func (r *medicine_repo) IsExistsByName(name string) bool {
	var count int

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, r.queryCountMedicine, name).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}

	return count > 0
}
