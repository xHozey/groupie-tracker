package main

import (
	"fmt"
	"log"
	"net/http"

	b "groupppietracker/handlers"
)

func main() {
	http.HandleFunc("/", b.HomePage)
	// http.HandleFunc("/", )
	fmt.Println("server runing in -> http://localhost:8484")
	err := http.ListenAndServe(":8484", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
