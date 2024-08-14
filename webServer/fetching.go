package groupie

import (
	"encoding/json"
	"log"
	"net/http"
)

func fetchIndex() []Artist {
	var artistians []Artist
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err1 := json.NewDecoder(resp.Body).Decode(&artistians)
	if err1 != nil {
		log.Fatal(err)
	}
	return artistians
}

func fetchArtist(s string) Result {
	var art Artist
	var loc Locations
	var dat Dates
	var rel Relation
	artist, errArt := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + s)
	if errArt != nil {
		log.Fatal(errArt)
	}
	defer artist.Body.Close()
	json.NewDecoder(artist.Body).Decode(&art)
	location, errLoc := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + s)
	if errLoc != nil {
		log.Fatal(errLoc)
	}
	json.NewDecoder(location.Body).Decode(&loc)
	defer location.Body.Close()
	date, errDate := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + s)
	if errDate != nil {
		log.Fatal(errDate)
	}
	json.NewDecoder(date.Body).Decode(&dat)
	defer date.Body.Close()
	relation, errRel := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + s)
	if errRel != nil {
		log.Fatal(errRel)
	}
	defer relation.Body.Close()
	json.NewDecoder(relation.Body).Decode(&rel)
	result := Result{art, loc, dat, rel}

	return result
}
