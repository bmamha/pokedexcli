package main

import (
	"fmt"
	"errors"
	"encoding/json"
  "github.com/bmamha/pokedexcli/internal/pokeapi"
)



func commandMapb(c *config) error {
     
	   if c.prevLocationsURL == nil {
			 return errors.New("you are on the first page")
		 }
		cachedURL := *c.prevLocationsURL
		if cachedData, ok := c.pokeClientCache.Get(cachedURL); !ok {
    locationsResp, err := c.pokeapiClient.GetLocations(c.prevLocationsURL)
    if err != nil {
			return err
		}

		jsonData, err := json.Marshal(locationsResp)
		if err != nil {
			return err 
		}

		c.pokeClientCache.Add(cachedURL, jsonData)

    c.nextLocationsURL = locationsResp.Next
		c.prevLocationsURL = locationsResp.Previous 

		for _, location := range locationsResp.Results {
			fmt.Println(location.Name)
    }
		} else {
			 var cachedResponse pokeapi.LocationArea 
		    err := json.Unmarshal(cachedData, &cachedResponse)
				if err != nil { 
					return err 
				}
		    for _, cachedLocation := range cachedResponse.Results {
			       fmt.Println(cachedLocation.Name)
		    } 
		}

		return nil

}
