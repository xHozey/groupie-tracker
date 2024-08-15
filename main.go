package main

import (
	"fmt"
	"log"
	"net/http"

	b "groupppietracker/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", b.HomePage)
	mux.HandleFunc("/artists/{id}", b.DetailsPage)
	// http.HandleFunc("/", )
	fmt.Println("server runing in -> http://localhost:8484")
	err := http.ListenAndServe(":8484", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
