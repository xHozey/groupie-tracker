package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// https://groupietrackers.herokuapp.com/api/artists/1
func DetailsPage(w http.ResponseWriter, r *http.Request) {
	var artist Artist
	pages := []string{
		"templates/artist_details.html",
		"templates/index.html",
	}
	tmpl, _ := template.ParseFiles(pages...)
	data2, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.NewDecoder(data2.Body).Decode(&artist)
	err_tmpl := tmpl.ExecuteTemplate(w, "base", artist)
	if err_tmpl != nil {
		fmt.Print(err_tmpl.Error())
		return
	}
}
