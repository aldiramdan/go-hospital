package diseases

import (
	"database/sql"
	"net/http"

	"github.com/aldiramdan/hospital/modules/medicines"
)

func DiseaseRoute(prefix string, db *sql.DB, mux *http.ServeMux) {

	repo_medicine := medicines.NewRepo(db)
	repo_disease := NewRepo(db)
	srvc := NewSrvc(repo_disease, repo_medicine)
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
			ctrl.GetByName(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc(prefix+"/convertcsv", ctrl.ConvertCSV)
	mux.HandleFunc(prefix+"/convertcsv/download", ctrl.Download)
}
