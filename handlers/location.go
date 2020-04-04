package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rmukubvu/amakhosi/model"
	"github.com/rmukubvu/amakhosi/repository"
	"io/ioutil"
	"net/http"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/location", add).Methods(http.MethodPost)
	api.HandleFunc("/location/{id}", search).Methods(http.MethodGet)
	return r
}

//PostLocation add location to database
func add(w http.ResponseWriter, req *http.Request) {
	//retrieve the json and unmarshal it to pumps
	var pump model.Pumps
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter valid data")
	}
	json.Unmarshal(reqBody, &pump)
	//here add the pump to the database
	err = repository.AddLocation(pump)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generateErrorMessage(err.Error()))
		return
	}
	//return valid status back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pump)
}

func search(w http.ResponseWriter, req *http.Request) {
	pathParams := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")

	if val, ok := pathParams["id"]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generateErrorMessage("id not specified or wrong id format"))
		return
	} else {
		res, _ := repository.LocationsById(val)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func generateErrorMessage(e string) []byte {
	ie := model.InternalError{Message: e}
	buf, err := json.Marshal(ie)
	if err != nil {
		return []byte(fmt.Sprintf(`{"message": "%s"}`, e))
	}
	return buf
}
