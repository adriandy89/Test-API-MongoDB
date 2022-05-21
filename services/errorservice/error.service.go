package errorservice

import (
	"encoding/json"
	"net/http"
)

func ErrorMessage(w http.ResponseWriter, message string, resStatus int) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(resStatus)
	json.NewEncoder(w).Encode(message)
}
