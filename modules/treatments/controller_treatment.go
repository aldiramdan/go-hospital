package treatments

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
	"github.com/aldiramdan/hospital/interfaces"
	"github.com/asaskevich/govalidator"
)

type treatment_ctrl struct {
	srvc interfaces.TreatmentSrvc
}

func NewCtrl(srvc interfaces.TreatmentSrvc) *treatment_ctrl {
	return &treatment_ctrl{srvc}
}

func (c *treatment_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodGet(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	c.srvc.GetAll().Send(w)
}

func (c *treatment_ctrl) GetById(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodGet(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	id := r.URL.Query().Get("id")

	c.srvc.GetById(id).Send(w)
}

func (c *treatment_ctrl) GetByNamePatient(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodGet(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	name := r.URL.Query().Get("name")

	c.srvc.GetByNamePatient(name).Send(w)
}

func (c *treatment_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodPost(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	var treatment models.Treatment
	err = json.NewDecoder(r.Body).Decode(&treatment)
	if err != nil {
		helpers.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&treatment)
	if err != nil {
		helpers.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.Add(treatment).Send(w)
}

func (c *treatment_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodPut(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.GetResponse(map[string]interface{}{"message": "Missing 'id' query parameter"}, 400, true).Send(w)
		return
	}

	var treatment models.Treatment
	err = json.NewDecoder(r.Body).Decode(&treatment)
	if err != nil {
		helpers.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&treatment)
	if err != nil {
		helpers.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.Update(id, treatment).Send(w)
}

func (c *treatment_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodDelete(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	id := r.URL.Query().Get("id")

	c.srvc.Delete(id).Send(w)
}

func (c *treatment_ctrl) ConvertCSV(w http.ResponseWriter, r *http.Request) {
	c.srvc.ConvertCSV().Send(w)
}

func (c *treatment_ctrl) Download(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(os.Getenv("PathtreatmentCSV"))
	if err != nil {
		helpers.GetResponse(err.Error(), 404, true).Send(w)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		helpers.GetResponse(err.Error(), 404, true).Send(w)
		return
	}

	c.srvc.Download().SendDownload(w, file, fileInfo)
}
