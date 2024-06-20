package routes

import (
	"encoding/json"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "GO is awesome"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
