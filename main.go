package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// type locations struct {
// 	Id       int      `json:"id"`
// 	Location []string `json:"locations"`
// }

// type dates struct {
// 	Id    int      `json:"id"`
// 	Dates []string `json:"dates"`
// }

//	type relation struct {
//		Id             int        `json:"id"`
//		DatesLocations [][]string `json:"datesLocations"`
//	}

var Art []artist

func main() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	err1 := json.NewDecoder(resp.Body).Decode(&Art)
	if err1 != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8070", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, Art)
}
