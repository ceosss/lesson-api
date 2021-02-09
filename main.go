package main

import (
	"fmt"
	"net/http"

	"github.com/ceosss/lesson-api/router"
)

func main() {
	fmt.Println("START")

	// Call router to get a newly configured router
	router := router.NewRouter()
	http.ListenAndServe(":3000", router)
}
