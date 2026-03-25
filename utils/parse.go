package utils

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`

	DatesLocations map[string][]string `json:"-"`
	Dates          []string            `json:"-"`
	Locations      []string            `json:"-"`
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

var Artistdata []Artist
var Relations Relation
var Locations Location
var Dates Date

// Fills the API's data to the variables.
func ApiParsing() error {
	artistrsp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil || artistrsp == nil {
		return err
	}

	relationrsp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil || relationrsp == nil {
		return err
	}

	locationrsp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil || locationrsp == nil {
		return err
	}

	datersp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil || datersp == nil {
		return err
	}

	defer artistrsp.Body.Close()
	defer relationrsp.Body.Close()
	defer locationrsp.Body.Close()
	defer datersp.Body.Close()

	err = json.NewDecoder(artistrsp.Body).Decode(&Artistdata)
	if err != nil {
		return err
	}

	err = json.NewDecoder(relationrsp.Body).Decode(&Relations)
	if err != nil {
		return err
	}

	err = json.NewDecoder(locationrsp.Body).Decode(&Locations)
	if err != nil {
		return err
	}

	err = json.NewDecoder(datersp.Body).Decode(&Dates)
	if err != nil {
		return err
	}

	for i := range Artistdata {
		for _, rel := range Relations.Index {
			if Artistdata[i].Id == rel.Id {
				Artistdata[i].DatesLocations = rel.DatesLocations
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
	return nil
}
