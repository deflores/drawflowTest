package drawflow

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetDrawflowsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	name := r.FormValue("name")

	var drs []DrawflowData
	var err error
	if len(name) > 0 {
		drs, err = GetDrawflows(name)
	} else {
		drs, err = GetAllDrawflows()
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(drs)
}
func GetDrawflowController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	drs, err := GetDrawflow(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(drs)
}

func SaveDrawflowsController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		return
	}
	var dr DrawflowData

	errParse := json.NewDecoder(r.Body).Decode(&dr)

	if errParse != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	drs, err := SaveDrawflow(dr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(drs)
}

func SavePythonCodeController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		return
	}

	var dr CodeDto

	errParse := json.NewDecoder(r.Body).Decode(&dr)

	if errParse != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err := WriteFile(dr.Code)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
		return
	}

}
