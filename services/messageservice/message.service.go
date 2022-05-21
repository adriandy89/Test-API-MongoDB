package messageservice

import (
	"Test-api-mongodb/models"
	"encoding/json"
	"net/http"
)

func SuccesMessage(w http.ResponseWriter, message string, resStatus int) {
	result := models.Response{
		Message: message,
	}
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resStatus)
	w.Write(response)
}
