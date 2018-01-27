package controller

import (
	"net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/juanmanuelromeraferrio/rest-api/model"
    "github.com/juanmanuelromeraferrio/rest-api/repository"
)

var people []string

func GetPeople(w http.ResponseWriter, r *http.Request) {
	people := repository.GetPeople()
    json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person := repository.GetPersonById(params["id"])
	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var data model.Person
	_ = json.NewDecoder(r.Body).Decode(&data)
	person := repository.CreatePerson(data)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person := repository.DeletePersonById(params["id"])
	json.NewEncoder(w).Encode(person)
}