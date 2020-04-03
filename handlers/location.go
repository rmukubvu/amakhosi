package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rmukubvu/amakhosi/model"
	"github.com/rmukubvu/amakhosi/repository"
	"io/ioutil"
	"net/http"
	"strconv"
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
	//retrieve the json and unmarshall it to pumps
	pump := model.Pumps{}
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter valid data")
	}
	json.Unmarshal(reqBody, &pump)
	//here add the pump to the database
	err = repository.AddLocation(pump)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	//id := req.URL.Query().Get("id")
	if val, ok := pathParams["id"]; ok {
		key, err := strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need an id"}`))
			return
		}
		res, _ := repository.LocationById(key)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
