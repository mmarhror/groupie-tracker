package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/functions"
	"groupie-tracker/tools"
)

func main() {
	artists := []tools.ArtistDetails{}
	functions.InitAllArtists(&artists)
	http.HandleFunc("/static/", functions.StyleFunc)
	http.HandleFunc("/", functions.GetAllArtists(&artists))
	http.HandleFunc("/artist/", functions.GetArtistDetails(&artists))
	http.HandleFunc("/filter", functions.FilterArtists(&artists))
	http.HandleFunc("/about", functions.AboutFunc)
	http.HandleFunc("/contact", functions.ContactFunc)
	fmt.Println("server is runing http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
