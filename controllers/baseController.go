package controllers

import "net/http"

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJson(w, code, map[string]string{"error": message})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {

}
