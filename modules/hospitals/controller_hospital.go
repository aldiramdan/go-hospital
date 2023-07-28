package hospitals

import (
	"net/http"

	"github.com/aldiramdan/hospital/helpers"
)

type (
	HospitalHandler struct{}
)

func (h *HospitalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"repository": "https://github.com/aldiramdan/go-hospitalnative",
		"demo":       "https://api-hospitalnative.aldep.site",
		"postman":    "https://documenter.getpostman.com/view/25646732/2s946pZUt7",
	}

	helpers.GetResponse(data, 200, false).Send(w)
}
