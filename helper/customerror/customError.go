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

// StatusBadRequest - Sets the error to Bad Request
func StatusBadRequest(response *http.ResponseWriter, err error) {
	fmt.Printf("ERROR: %+v", err)
	(*response).WriteHeader(http.StatusBadRequest)
	(*response).Write([]byte(`{"message": "` + err.Error() + `"}`))
}
