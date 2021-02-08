package main

import (
	"fmt"
	"net/http"

	"github.com/ceosss/lesson-api/router"
)

func main() {
	fmt.Println("START")
	router := router.NewRouter()
	http.ListenAndServe(":3000", router)
}
