package main

import (
	
	"fmt"
 
  "github.com/bmamha/pokedexcli/internal/pokeapi"
	
	"encoding/json"
)



func commandMap(c *config) error {
	var cachedURL string
	if c.nextLocationsURL == nil {
     cachedURL = "fpage" 
   } else {
	    cachedURL = *c.nextLocationsURL
} 
	if cachedData, ok := c.pokeClientCache.Get(cachedURL); !ok {
    locationsResp, err := c.pokeapiClient.GetLocations(c.nextLocationsURL)	
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




