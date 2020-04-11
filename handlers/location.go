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

type MarkerState struct {
	Index int `json:"id"`
}

var marker MarkerState

func InitRouter() *mux.Router {
	marker.Index = 20
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/location", add).Methods(http.MethodPost)
	api.HandleFunc("/location/{id}", search).Methods(http.MethodGet)
	api.HandleFunc("/test", test).Methods(http.MethodGet)
	return r
}

func test(w http.ResponseWriter, req *http.Request) {
	marker.Index++
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(marker)
}

//PostLocation add location to database
func add(w http.ResponseWriter, req *http.Request) {
	//retrieve the json and unmarshal it to pumps
	pump := model.Pumps{}
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(generateErrorMessage(err.Error()))
		return
	}
	//check if its a valid json string
	if ok := validJson(string(reqBody)); !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(generateErrorMessage("invalid json string"))
		return
	}
	//continue to unmarshall
	json.Unmarshal(reqBody, &pump)
	//here add the pump to the database
	err = repository.AddLocation(pump)
	//set the header
	//return valid status back
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(generateErrorMessage(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pump)
}

func search(w http.ResponseWriter, req *http.Request) {
	pathParams := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")

	if val, ok := pathParams["id"]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(generateErrorMessage("id not specified or wrong id format")))
		return
	} else {
		res, _ := repository.LocationsById(val)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func validJson(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func generateErrorMessage(e string) string {
	ie := model.InternalError{Message: e}
	buf, err := json.Marshal(ie)
	if err != nil {
		return string([]byte(fmt.Sprintf(`{"message": "%s"}`, e)))
	}
	return string(buf)
}
