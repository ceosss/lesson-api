package router

import (
	"github.com/ceosss/lesson-api/handlers/modelhandlers"
	"github.com/gorilla/mux"
)

//NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/model", modelhandlers.AllModels).Methods("GET")
	router.HandleFunc("/model", modelhandlers.CreateModel).Methods("POST")
	router.HandleFunc("/model/{id}", modelhandlers.SingleModel).Methods("GET")
	router.HandleFunc("/model/{id}", modelhandlers.DeleteModel).Methods("DELETE")
	return router
}
