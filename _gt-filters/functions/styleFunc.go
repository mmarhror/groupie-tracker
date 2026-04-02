package functions

import (
	"net/http"
	"os"
)

func StyleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[1:]
	File, err := os.Stat(filePath)
	if err != nil {
		HandleError(w, "Not Found!", http.StatusNotFound)
		return
	}
	if File.IsDir() {
		HandleError(w, "Forbidden!", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, filePath)
}
