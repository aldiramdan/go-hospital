package hospitals

import "net/http"

func HospitalRoute(prefix string, mux *http.ServeMux) {
	mux.Handle(prefix+"/", &HospitalHandler{})
}
