package successresponse

import "net/http"

func setContentType(response *http.ResponseWriter) {
	(*response).Header().Set("content-type", "application/json")

}

// Created ...
func Created(response *http.ResponseWriter) {
	setContentType(response)
	(*response).WriteHeader(201)
}

// OK ...
func OK(response *http.ResponseWriter) {
	setContentType(response)
	(*response).WriteHeader(200)
}

// NoContent ...
func NoContent(response *http.ResponseWriter) {
	setContentType(response)
	(*response).WriteHeader(204)
}
