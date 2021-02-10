package customerror

import (
	"fmt"
	"net/http"
)

// InternalServerError - Sets the error to the response
func InternalServerError(response *http.ResponseWriter, err error) {
	fmt.Printf("ERROR: %v", err)
	(*response).WriteHeader(http.StatusInternalServerError)
	(*response).Write([]byte(`{"message": "` + err.Error() + `"}`))
}
