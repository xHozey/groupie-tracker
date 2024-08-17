package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	var artists []Artist
	pages := []string{
		"templates/artist_details.html",
		"templates/index.html",
	}
	tmpl, err := template.ParseFiles(pages...)
	if err != nil {
		log.Fatal(err)
	}
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	err = json.NewDecoder(data.Body).Decode(&artists)
	if err != nil {
		log.Fatal(err)
	}
	err_tmpl := tmpl.ExecuteTemplate(w, "base", artists)
	if err_tmpl != nil {
		fmt.Print(err_tmpl.Error())
		return
	}
	tmpl.Execute(w, artists)
}
