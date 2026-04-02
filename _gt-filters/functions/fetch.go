package functions

import (
	"encoding/json"
	"net/http"
)

func Fetch(url string, data interface{}) error {
	response, errGet := http.Get(url)
	if errGet != nil {
		return errGet
	}
	defer response.Body.Close()
	errDecode := json.NewDecoder(response.Body).Decode(data)
	if errDecode != nil {
		return errDecode
	}
	return nil
}
