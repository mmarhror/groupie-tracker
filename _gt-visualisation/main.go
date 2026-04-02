package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/helpers"
)

func main() {
	http.HandleFunc("/media/", helpers.MediaHandler)

	http.HandleFunc("/templates/", helpers.CssHandler)

	http.HandleFunc("/", helpers.HomeHandler)
	
	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
