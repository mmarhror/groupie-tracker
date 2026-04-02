package functions

import (
	"fmt"
	"groupie-tracker/tools"
	"log"
	"strconv"
)

func InitAllArtists(artists *[]tools.ArtistDetails) {
	fmt.Println("Fetching all data from API... Please wait.")

	var basicArtists []tools.Artists
	err := Fetch("https://groupietrackers.herokuapp.com/api/artists", &basicArtists)
	if err != nil {
		log.Fatalln("Error fetching basic artists:", err)
	}

	for _, art := range basicArtists {
		details, err := FetchAllData(strconv.Itoa(art.Id))
		if err != nil {
			log.Println("Error fetching details for artist ID", art.Id)
			continue
		}
		*artists = append(*artists, details)
	}

	fmt.Println("Data fetched successfully!")
}
