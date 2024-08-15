package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	var artists []Artist
	pages := []string{
		"templates/index.html",
	}
	tmpl, _ := template.ParseFiles(pages...)
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	json.NewDecoder(data.Body).Decode(&artists)
	err_tmpl := tmpl.ExecuteTemplate(w, "base", artists)
	if err_tmpl != nil {
		fmt.Print(err_tmpl.Error())
		return
	}
}
