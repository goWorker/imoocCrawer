package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Moive struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Moive{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphery Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jack"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON Marshal failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON Unmarshal failed: %s", err)
	}
	fmt.Println(titles)
}
