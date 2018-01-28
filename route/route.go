package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juanmanuelromeraferrio/rest-api/controller"
)

func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/people", controller.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", controller.GetPerson).Methods("GET")
	router.HandleFunc("/people", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", controller.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
