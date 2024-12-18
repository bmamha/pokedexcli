package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)



func commandMap(c *config) error {
	
	url := c.Next 
	if url != nil {
	res, err := http.Get(*url)
     if err != nil {
			 log.Fatal(err)
		 }

		 var locations Location

		 defer res.Body.Close()

		 decoder := json.NewDecoder(res.Body)
		 if err := decoder.Decode(&locations); err != nil {
			 log.Fatal(err)
		 } 
		 
		for _, location := range locations.Results {
			fmt.Printf("%v\n", location.Name)

		}

		next_url := locations.Next 
		c.Next = &next_url
		prev_url := locations.Previous
		c.Previous = &prev_url 
		} else {
			fmt.Println("You are on the last page")
		}
		return nil
}




type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

