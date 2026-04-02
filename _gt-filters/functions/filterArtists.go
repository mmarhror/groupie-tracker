package functions

import (
	"fmt"
	"groupie-tracker/tools"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func FilterArtists(artists *[]tools.ArtistDetails) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		r.ParseForm()
		creationDate, ok := r.PostForm["creation-date"]
		if !ok {
			HandleError(w, "Status Bad Request", http.StatusBadRequest)
			return
		}
		creation_date, err := strconv.Atoi(creationDate[0])
		if err != nil {
			HandleError(w, "Status Bad Request", http.StatusBadRequest)
			return
		}
		firstAlbumDate, ok := r.PostForm["first-album-date"]
		if !ok {
			HandleError(w, "Status Bad Request", http.StatusBadRequest)
			return
		}
		firstAlbumStr := firstAlbumDate[0]
		fmt.Println(firstAlbumStr)
		if len(firstAlbumStr) < 4 {
			HandleError(w, "Status Bad Request", http.StatusBadRequest)
			return
		}
		first_Album_Date, err := strconv.Atoi(firstAlbumStr[len(firstAlbumStr)-4:])
		if err != nil {
			HandleError(w, "Status Bad Request", http.StatusBadRequest)
			return
		}
		locations_Search, ok := r.PostForm["locations"]
		if !ok {
			HandleError(w, "Status Bad Request", http.StatusBadRequest)
			return
		}

		members := r.PostForm["members"]
		oneTree := false
		fourSix := false
		plusSeven := false
		for _, val := range members {
			if val == "1-3" {
				oneTree = true
				continue
			}
			if val == "4-6" {
				fourSix = true
				continue
			}
			if val == "plus7" {
				plusSeven = true
				continue
			}
		}
		minMembers, maxMembers := GetMinMaxMembers(oneTree, fourSix, plusSeven)
		var artistDetails []tools.ArtistDetails
		for _, artsDets := range *artists {
			matchLocation := false
			matchFilters := false
			if len(locations_Search[0]) > 0 {
				for _, location := range artsDets.Locations {
					if strings.Contains(location, locations_Search[0]) {
						matchLocation = true
					}
				}
			}
			FirstAlbumYear, err := strconv.Atoi(artsDets.FirstAlbum[len(artsDets.FirstAlbum)-4:])
			if err != nil {
				HandleError(w, "Status Bad Request", http.StatusBadRequest)
				return
			}
			if artsDets.CreationDate <= creation_date || FirstAlbumYear <= first_Album_Date || (len(artsDets.Members) >= minMembers && len(artsDets.Members) <= maxMembers) {
				matchFilters = true
			}
			if matchLocation || matchFilters {
				artistDetails = append(artistDetails, artsDets)
			}
		}

		temp, errParse := template.ParseFiles("templates/filter.html")
		if errParse != nil {
			HandleError(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, artistDetails)
	}
}
