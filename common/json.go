package common

import (
	"encoding/json"
	"net/http"
)

func MarshalJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func UnmarshalJSON(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteErrorJSON(w http.ResponseWriter, status int, message string) {
	MarshalJSON(w, status, map[string]string{"error": message})
}
