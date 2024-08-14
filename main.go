package main

import (
	"log"
	"net/http"

	web "groupie/webServer"
)

func main() {
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/artist", web.ArtistInfo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
