package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Artist struct {
	Id           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Concerts     map[string][]string `json:"-"`
	Dates        []string            `json:"-"`
	Locations    []string            `json:"-"`
}

type Relation struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type Date struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Location struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

var (
	Artistdata []Artist
	Relations  Relation
	Locations  Location
	Dates      Date
)

// Takes an array of errors and checks if any of them is not nil .
func NoError(errs []error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func StatCode(states []*http.Response) int {
	for _, stat := range states {
		if stat == nil {
			return 500
		}
		if stat.StatusCode != 200 {
			return stat.StatusCode
		}
	}
	return 200
}

// Fills the API's data to the variables.
func ApiParsing() (int, error) {
	artistrsp, err1 := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	relationrsp, err2 := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	locationrsp, err3 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	datersp, err4 := http.Get("https://groupietrackers.herokuapp.com/api/dates")

	errs := []error{err1, err2, err3, err4}
	stat := []*http.Response{artistrsp, relationrsp, locationrsp, datersp}

	if StatCode(stat) >= 400 || NoError(errs) != nil {
		return StatCode(stat), errors.New("error fetching data")
	}

	defer artistrsp.Body.Close()
	defer relationrsp.Body.Close()
	defer locationrsp.Body.Close()
	defer datersp.Body.Close()

	errs = append(errs, json.NewDecoder(artistrsp.Body).Decode(&Artistdata))
	errs = append(errs, json.NewDecoder(relationrsp.Body).Decode(&Relations))
	errs = append(errs, json.NewDecoder(locationrsp.Body).Decode(&Locations))
	errs = append(errs, json.NewDecoder(datersp.Body).Decode(&Dates))

	if NoError(errs) != nil {
		fmt.Println(NoError(errs))
		return 500, errors.New("failed to decode /api")
	}

	for i := range Artistdata {
		for _, rel := range Relations.Index {
			if Artistdata[i].Id == rel.Id {
				Artistdata[i].Concerts = rel.DatesLocations
				break
			}
		}
		for _, d := range Dates.Index {
			if Artistdata[i].Id == d.Id {
				Artistdata[i].Dates = d.Dates
				break
			}
		}
		for _, lo := range Locations.Index {
			if Artistdata[i].Id == lo.Id {
				Artistdata[i].Locations = lo.Locations
				break
			}
		}
	}
	return 200, nil
}
