package diseases

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aldiramdan/hospital/databases/db/models"
	"github.com/aldiramdan/hospital/helpers"
	"github.com/aldiramdan/hospital/interfaces"
	"github.com/asaskevich/govalidator"
)

type disease_ctrl struct {
	srvc interfaces.DiseaseSrvc
}

func NewCtrl(srvc interfaces.DiseaseSrvc) *disease_ctrl {
	return &disease_ctrl{srvc}
}

func (c *disease_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodGet(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	c.srvc.GetAll().Send(w)
}

func (c *disease_ctrl) GetById(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodGet(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	id := r.URL.Query().Get("id")

	c.srvc.GetById(id).Send(w)
}

func (c *disease_ctrl) GetByName(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodGet(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	name := r.URL.Query().Get("name")

	c.srvc.GetByName(name).Send(w)
}

func (c *disease_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodPost(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	var disease models.Disease
	err = json.NewDecoder(r.Body).Decode(&disease)
	if err != nil {
		helpers.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&disease)
	if err != nil {
		helpers.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.Add(disease).Send(w)
}

func (c *disease_ctrl) Update(w http.ResponseWriter, r *http.Request) {
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

	var disease models.Disease
	err = json.NewDecoder(r.Body).Decode(&disease)
	if err != nil {
		helpers.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&disease)
	if err != nil {
		helpers.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.Update(id, disease).Send(w)
}

func (c *disease_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	err := helpers.MethodDelete(w, r)
	if err != nil {
		helpers.GetResponse(err.Error(), 405, true).Send(w)
		return
	}

	id := r.URL.Query().Get("id")

	c.srvc.Delete(id).Send(w)
}

func (c *disease_ctrl) ConvertCSV(w http.ResponseWriter, r *http.Request) {
	c.srvc.ConvertCSV().Send(w)
}

func (c *disease_ctrl) Download(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(os.Getenv("PathDiseaseCSV"))
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
