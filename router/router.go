package router

import (
	"github.com/ceosss/lesson-api/handlers/modelhandlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/model", modelhandlers.AllModels).Methods("GET")
	router.HandleFunc("/model", modelhandlers.CreateModel).Methods("POST")
	return router
}
