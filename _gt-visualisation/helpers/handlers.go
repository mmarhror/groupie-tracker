package helpers

import (
	"errors"
	"net/http"
	"text/template"
)

// ErrorData holds the structure for rendering custom error pages.
// Message to display to the user
// HTTP status code (e.g. 404, 500)
type ErrorData struct {
	ErrorMessage string
	ErrorNumber  int
}

// Global variables to store the error template and any parsing error.
var (
	tmp      *template.Template
	templErr error
)

func init() {
	tmp, templErr = template.ParseFiles("templates/errors.html")
}

// InitErr writes a custom error page to the response writer.
// It sets the status code, prepares the data, and executes the error template.
func InitErr(w http.ResponseWriter, e error, n int) {
	if templErr != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(n)
	errData := ErrorData{
		ErrorMessage: e.Error(),
		ErrorNumber:  n,
	}

	tmp.Execute(w, errData)
}

// HomeHandler handles requests to the root path ("/").
// If the path is anything else, it returns a 404 error.
// Otherwise, it fetches data from APIs and renders the main artist page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		err := errors.New("method not allowed")
		InitErr(w, err, 405)
		return
	}
	if r.URL.Path != "/" {
		err := errors.New("path Not found")
		InitErr(w, err, 404)
		return
	}
	nmber, err := ApiParsing()
	if err != nil {
		InitErr(w, err, nmber)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		err = errors.New("template error")
		InitErr(w, err, 500)
		return
	}

	tmpl.Execute(w, Artistdata)
}

// CssHandler serves the CSS file located at "templates/style.css".
// If the requested URL path is anything other than "/templates/style.css",
// it returns a 404 Not Found error using InitErr.
func CssHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("C   " + r.URL.Path)
	if r.URL.Path != "/templates/style.css" {
		err := errors.New("not found")
		InitErr(w, err, 404)
		return
	}
	http.ServeFile(w, r, "templates/style.css")
}

// MediaHandler serves the image file "media/undo.png".
// If the requested URL path is anything other than "/media/undo.png",
// it returns a 404 Not Found error using InitErr.
func MediaHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("M   " + r.URL.Path)
	if r.URL.Path != "/media/undo.png" {
		err := errors.New("not found")
		InitErr(w, err, 404)
		return
	}
	http.ServeFile(w, r, "media/undo.png")
}
