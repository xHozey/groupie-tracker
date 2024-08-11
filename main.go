package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Define a struct to match the expected JSON structure.
type Response struct {
	Artistsinfo Artists `json:"artists"`
	Locations   string  `json:"locations"`
	Dates       string  `json:"dates"`
	Relation    string  `json:"relation"`
}

type Artists struct {
	id           int      `json.id`
	image        string   `json.image`
	name         string   `json.name`
	members      []string `json.members`
	creationDate int      `json.creationDate`
	firstAlbum   string   `json.firstAlbum`
}

func main() {
	// Perform the HTTP GET request.
	respondData, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		os.Exit(1)
	}
	defer respondData.Body.Close() // Ensure the body is closed.

	// Read the response body.
	data, err := io.ReadAll(respondData.Body)
	if err != nil {
		fmt.Println("Error reading data:", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data into a Response struct.
	var responseObject Artists
	err = json.Unmarshal(data, &responseObject)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	// Print the parsed data.

	fmt.Print(responseObject.name)
}
