package controllers

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	RespondWithError(w http.ResponseWriter, code int, message string)
	RespondWithJson(w http.ResponseWriter, code int, payload interface{})
}

/**
* Error Response
 */
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJson(w, code, map[string]string{"error": message})
}

/**
* Write Out Json
 */
func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
