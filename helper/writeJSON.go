package helper

import (
	"encoding/json"
	"foodieland/model/web"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(web.WebResponse{
		Code:   code,
		Status: "OK",
		Data:   data,
	})
}

func WriteErrorJSON(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(web.ErrorResponse{
		Code:   code,
		Message: message,
	})
}
