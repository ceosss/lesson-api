package router

import (
	"github.com/ceosss/lesson-api/handlers/lessonhandlers"
	"github.com/ceosss/lesson-api/handlers/modelhandlers"
	"github.com/gorilla/mux"
)

//NewRouter - Return a new handler with configured routes
func NewRouter() *mux.Router {

	// New Multiplexer
	router := mux.NewRouter()

	// All the handlers

	// Model Handlers
	router.HandleFunc("/model", modelhandlers.AllModels).Methods("GET")
	router.HandleFunc("/model", modelhandlers.CreateModel).Methods("POST")
	router.HandleFunc("/model/{id}", modelhandlers.SingleModel).Methods("GET")
	router.HandleFunc("/model/{id}", modelhandlers.UpdateModel).Methods("PUT")
	router.HandleFunc("/model/{id}", modelhandlers.DeleteModel).Methods("DELETE")

	// Lesson Handlers
	router.HandleFunc("/lesson", lessonhandlers.AllLessons).Methods("GET")
	router.HandleFunc("/lesson", lessonhandlers.CreateLesson).Methods("POST")
	router.HandleFunc("/lesson/{id}", lessonhandlers.SingleLesson).Methods("GET")
	router.HandleFunc("/lesson/{id}", lessonhandlers.UpdateLesson).Methods("PUT")
	router.HandleFunc("/lesson/{id}", lessonhandlers.DeleteLesson).Methods("DELETE")
	return router
}
