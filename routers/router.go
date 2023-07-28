package routers

import (
	"database/sql"
	"net/http"

	"github.com/aldiramdan/hospital/modules/diseases"
	"github.com/aldiramdan/hospital/modules/doctors"
	"github.com/aldiramdan/hospital/modules/hospitals"
	"github.com/aldiramdan/hospital/modules/medicines"
	"github.com/aldiramdan/hospital/modules/patients"
	"github.com/aldiramdan/hospital/modules/treatments"
)

func IndexRoute(mux *http.ServeMux, db *sql.DB) {
	hospitals.HospitalRoute("", mux)
	patients.PatientRoute("/patient", db, mux)
	doctors.DoctorRoute("/doctor", db, mux)
	medicines.MedicineRoute("/medicine", db, mux)
	diseases.DiseaseRoute("/disease", db, mux)
	treatments.TreatmentRoute("/treatment", db, mux)
}
