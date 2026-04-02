package functions

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/tools"
)

func GetArtistDetails(artists *[]tools.ArtistDetails) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		temp, errParse := template.ParseFiles("templates/artist.html")
		if errParse != nil {
			HandleError(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		path := r.URL.Path
		id := path[len("/artist/"):]

		if id == "" {
			HandleError(w, "Bad Request: Missing ID", http.StatusBadRequest)
			return
		}
		found := false
		var artDetails tools.ArtistDetails
		for _, artist := range *artists {
			if strconv.Itoa(artist.Id) == id {
				artDetails = artist
				found = true
			}
		}
		if !found {
			HandleError(w, fmt.Sprintf("Artist with ID %s not found", id), http.StatusNotFound)
			return
		}
		err2 := temp.Execute(w, artDetails)
		if err2 != nil {
			HandleError(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
