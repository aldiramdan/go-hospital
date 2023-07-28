package diseases

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aldiramdan/hospital/databases/db/models"
)

type disease_repo struct {
	queryGetAllDisease    string
	queryGetByIdDisease   string
	queryGetByNameDisease string
	queryAddDisease       string
	queryUpdateDisease    string
	queryDeleteDisease    string
	queryCountDisease     string
	db                    *sql.DB
}

func NewRepo(db *sql.DB) *disease_repo {
	return &disease_repo{
		queryGetAllDisease: `SELECT * FROM disease AS d 
								JOIN medicine AS m ON d.medicine_id = m.id`,
		queryGetByIdDisease: `SELECT * FROM disease AS d 
								JOIN medicine AS m ON d.medicine_id = m.id 
								WHERE d.id=?`,
		queryGetByNameDisease: `SELECT * FROM disease AS d 
									JOIN medicine AS m ON d.medicine_id = m.id 
									WHERE d.name=?`,
		queryAddDisease:    "INSERT INTO disease(id, name, medicine_id) VALUES(?,?,?)",
		queryUpdateDisease: "UPDATE disease SET name=?, medicine_id=? WHERE id=?",
		queryDeleteDisease: "DELETE FROM disease WHERE id=?",
		queryCountDisease:  "SELECT COUNT(*) FROM disease WHERE name=?",
		db:                 db,
	}
}

func (r *disease_repo) GetAll() (*models.Diseases, error) {
	var disease models.Disease
	var diseases models.Diseases

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, r.queryGetAllDisease)
	if err != nil {
		return nil, errors.New("failed get data disease")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&disease.DiseaseId, &disease.Name, &disease.Medicine.MedicineId, &disease.CreatedAt, &disease.UpdatedAt, &disease.Medicine.MedicineId, &disease.Medicine.Name, &disease.Medicine.Price, &disease.Medicine.CreatedAt, &disease.Medicine.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data disease")
		}
		diseases = append(diseases, disease)
	}

	return &diseases, nil
}

func (r *disease_repo) GetById(id string) (*models.Disease, error) {
	var disease models.Disease

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows := r.db.QueryRowContext(ctx, r.queryGetByIdDisease, id)
	err := rows.Scan(&disease.DiseaseId, &disease.Name, &disease.Medicine.MedicineId, &disease.CreatedAt, &disease.UpdatedAt, &disease.Medicine.MedicineId, &disease.Medicine.Name, &disease.Medicine.Price, &disease.Medicine.CreatedAt, &disease.Medicine.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("data disease not found")
	} else if err != nil {
		return nil, errors.New("failed to scan data disease")
	}

	return &disease, nil
}

func (r *disease_repo) GetByName(name string) (*models.Disease, error) {
	var disease models.Disease

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows := r.db.QueryRowContext(ctx, r.queryGetByNameDisease, name)
	err := rows.Scan(&disease.DiseaseId, &disease.Name, &disease.Medicine.MedicineId, &disease.CreatedAt, &disease.UpdatedAt, &disease.Medicine.MedicineId, &disease.Medicine.Name, &disease.Medicine.Price, &disease.Medicine.CreatedAt, &disease.Medicine.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("data disease not found")
	} else if err != nil {
		return nil, errors.New("failed to scan data disease")
	}

	return &disease, nil
}

func (r *disease_repo) Add(data models.Disease) (*models.Disease, error) {
	if exists := r.IsExistsByName(data.Name); exists {
		return nil, errors.New("name disease has been exists")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, r.queryAddDisease, data.DiseaseId, data.Name, data.Medicine.MedicineId)
	if err != nil {
		return nil, errors.New("failed to add data disease")
	}

	result, err := r.GetById(data.DiseaseId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *disease_repo) Update(id string, data models.Disease) (*models.Disease, error) {
	if exists := r.IsExistsByName(data.Name); exists {
		return nil, errors.New("name disease has been exists")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, r.queryUpdateDisease, data.Name, data.Medicine.MedicineId, data.DiseaseId)
	if err != nil {
		return nil, errors.New("failed to update data disease")
	}

	result, err := r.GetById(data.DiseaseId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *disease_repo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.QueryContext(ctx, r.queryDeleteDisease, id)
	if err != nil {
		return errors.New("failed get data disease")
	}

	return nil
}

func (r *disease_repo) IsExistsByName(name string) bool {
	var count int

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, r.queryCountDisease, name).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}

	return count > 0
}

func (r *disease_repo) DuplicateName(data models.Diseases) (bool, error) {
	diseaseMap := make(map[string]bool)

	for _, d := range data {
		name := d.Name
		if diseaseMap[name] {
			return true, errors.New("data disease duplicate")
		}
		diseaseMap[name] = true
	}

	return false, nil
}
