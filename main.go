package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("START")
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	http.ListenAndServe(":3000", router)
}
func home(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("LESSON API"))
}
