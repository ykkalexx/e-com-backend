package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ParseJSON parses a JSON request and decodes it into the provided payload
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

// WriteJSON writes a JSON response to the client
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

/* 
 WriteError writes an error response to the client
 We can use this function to write error responses to the client in a consistent format.
*/
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}