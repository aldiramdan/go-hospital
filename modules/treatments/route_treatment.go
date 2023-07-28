package treatments

import (
	"database/sql"
	"net/http"

	"github.com/aldiramdan/hospital/modules/diseases"
	"github.com/aldiramdan/hospital/modules/doctors"
	"github.com/aldiramdan/hospital/modules/patients"
)

func TreatmentRoute(prefix string, db *sql.DB, mux *http.ServeMux) {

	repo_patient := patients.NewRepo(db)
	repo_disease := diseases.NewRepo(db)
	repo_doctor := doctors.NewRepo(db)
	repo_treatment := NewRepo(db)
	srvc := NewSrvc(repo_patient, repo_disease, repo_doctor, repo_treatment)
	ctrl := NewCtrl(srvc)

	mux.HandleFunc(prefix+"/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ctrl.GetAll(w, r)
		case http.MethodPost:
			ctrl.Add(w, r)
		case http.MethodPut:
			ctrl.Update(w, r)
		case http.MethodDelete:
			ctrl.Delete(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc(prefix+"/q", func(w http.ResponseWriter, r *http.Request) {
		queryID := r.URL.Query().Get("id")
		queryName := r.URL.Query().Get("name")

		switch {
		case queryID != "":
			ctrl.GetById(w, r)
		case queryName != "":
			ctrl.GetByNamePatient(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc(prefix+"/convertcsv", ctrl.ConvertCSV)
	mux.HandleFunc(prefix+"/convertcsv/download", ctrl.Download)
}
