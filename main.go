package main

import (
	"fmt"
	"groupie-tracker/utils"
	"net/http"
)

func main() {

	err := utils.ApiParsing()
	if err != nil {
		fmt.Println(err)
	}

	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))

	http.HandleFunc("/template/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "template/style.css")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := utils.RenderPage(w, "template/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	fmt.Println("Running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
