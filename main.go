package main

import (
	"log"
	"net/http"
	"task-5-pbi-fullstack-developer-adityap/controllers"
	"task-5-pbi-fullstack-developer-adityap/models"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectionDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/photos", controllers.Index).Methods("GET")
	r.HandleFunc("/photo/{id}", controllers.Show).Methods("GET")
	r.HandleFunc("/photo", controllers.Create).Methods("POST")
	r.HandleFunc("/photo/{id}", controllers.Update).Methods("PUT")
	r.HandleFunc("/photo", controllers.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
