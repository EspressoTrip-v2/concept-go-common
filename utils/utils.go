package utils

import (
	"encoding/json"
	"fmt"
	libErrors "github.com/EspressoTrip-v2/concept-go-common/liberrors"
	"net/http"
)

// WriteResponse is a convenience method for json response replies in Gorilla mux
// it adds the requited Content-Type and header code
func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("[response:writer]: Failed to convert data to JSON, could not send")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(libErrors.NewBadRequestError(err.Error()))
	}
}
