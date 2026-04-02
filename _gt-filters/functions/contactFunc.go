package functions

import (
	"html/template"
	"net/http"
)

func ContactFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("templates/contact.html")
	if err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	execError := temp.Execute(w, nil)
	if execError != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
