package router

import (
	"github.com/ceosss/lesson-api/handlers/lessonhandlers"
	"github.com/ceosss/lesson-api/handlers/modelhandlers"
	"github.com/gorilla/mux"
)

//NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/model", modelhandlers.AllModels).Methods("GET")
	router.HandleFunc("/model", modelhandlers.CreateModel).Methods("POST")
	router.HandleFunc("/model/{id}", modelhandlers.SingleModel).Methods("GET")
	router.HandleFunc("/model/{id}", modelhandlers.UpdateModel).Methods("PUT")
	router.HandleFunc("/model/{id}", modelhandlers.DeleteModel).Methods("DELETE")
	router.HandleFunc("/lesson", lessonhandlers.AllLessons).Methods("GET")
	router.HandleFunc("/lesson", lessonhandlers.CreateLesson).Methods("POST")
	router.HandleFunc("/lesson/{id}", lessonhandlers.SingleLesson).Methods("GET")
	return router
}
