package main

import (
	"fmt"
	"net/http"

	"github.com/ceosss/lesson-api/handlers/modelhandlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("START")
	router := mux.NewRouter()
	router.HandleFunc("/model", modelhandlers.AllModels).Methods("GET")
	router.HandleFunc("/model", modelhandlers.CreateModel).Methods("POST")
	http.ListenAndServe(":3000", router)
}
