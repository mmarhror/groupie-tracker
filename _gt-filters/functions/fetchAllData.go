package functions

import (
	"fmt"

	"groupie-tracker/tools"
)


func FetchAllData(id string) (tools.ArtistDetails, error) {
	var artistDetails tools.ArtistDetails
	urlArtist := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id)
	var artist tools.Artists
	errArtistFetch := Fetch(urlArtist, &artist)
	if errArtistFetch != nil {
		return tools.ArtistDetails{}, errArtistFetch
	}
	var locations tools.Locations
	errLocationsFetch := Fetch(artist.Locations, &locations)
	if errLocationsFetch != nil {
		return tools.ArtistDetails{}, errLocationsFetch
	}
	var dates tools.Dates
	errDatesFetch := Fetch(artist.ConcertDates, &dates)
	if errDatesFetch != nil {
		return tools.ArtistDetails{}, errDatesFetch
	}
	var relations tools.Relations
	errRelationsFetch := Fetch(artist.Relations, &relations)
	if errRelationsFetch != nil {
		return tools.ArtistDetails{}, errRelationsFetch
	}

	artistDetails.Id = artist.Id
	artistDetails.Image = artist.Image
	artistDetails.Name = artist.Name
	artistDetails.Members = artist.Members
	artistDetails.CreationDate = artist.CreationDate
	artistDetails.FirstAlbum = artist.FirstAlbum
	artistDetails.Locations = locations.Locations
	artistDetails.Dates = dates.Dates
	artistDetails.DatesLocations = relations.DatesLocations

	return artistDetails, nil
}
