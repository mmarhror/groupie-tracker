package functions

import (
	"groupie-tracker/tools"
	"html/template"
	"net/http"
)

func GetAllArtists(artists *[]tools.ArtistDetails) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			HandleError(w, "Not Found", http.StatusNotFound)
			return
		}
		if r.Method != http.MethodGet {
			HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		temp, errParse := template.ParseFiles("templates/index.html")
		if errParse != nil {
			HandleError(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, *artists)
	}
}
