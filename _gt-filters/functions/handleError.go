package functions

import (
	"html/template"
	"net/http"
	"strconv"
)

func HandleError(w http.ResponseWriter, errorText string, statusCode int) {
	errorData := make(map[string]string)
	errorData["errorText"] = errorText
	errorData["statusCode"] = strconv.Itoa(statusCode)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error!", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	execError := tmpl.Execute(w, errorData)
	if execError != nil {
		http.Error(w, "500 Internal Server Error!", http.StatusInternalServerError)
		return
	}
}
