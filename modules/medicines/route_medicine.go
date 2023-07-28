package medicines

import (
	"database/sql"
	"net/http"
)

func MedicineRoute(prefix string, db *sql.DB, mux *http.ServeMux) {

	repo := NewRepo(db)
	srvc := NewSrvc(repo)
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
