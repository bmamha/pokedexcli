package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)



func commandMapb(c *config) error {

	   url := c.Previous
    
if 	*url != "" {
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
		c.Previous = &prev_url} else {
			fmt.Printf("You are on the first page!\n")
		}
		return nil

}
