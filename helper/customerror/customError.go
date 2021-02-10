package customerror

import (
	"fmt"
	"net/http"
)

// InternalServerError - Sets the error to Internal Server Error
func InternalServerError(response *http.ResponseWriter, err error) {
	fmt.Printf("ERROR: %v", err)
	(*response).WriteHeader(http.StatusInternalServerError)
	(*response).Write([]byte(`{"message": "` + err.Error() + `"}`))
}

// BadRequest - Sets the error to Bad Request
func BadRequest(response *http.ResponseWriter, err error) {
	fmt.Printf("ERROR: %+v", err)
	(*response).WriteHeader(http.StatusBadRequest)
	(*response).Write([]byte(`{"message": "` + err.Error() + `"}`))
}

// Unauthorized - Sets the error to Unauthorized
func Unauthorized(response *http.ResponseWriter, err error) {
	fmt.Printf("ERROR: %+v", err)
	(*response).WriteHeader(http.StatusUnauthorized)
	(*response).Write([]byte(`{"message": "` + err.Error() + `"}`))
}
