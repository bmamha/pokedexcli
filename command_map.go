package main

import (
	
	"fmt"
)



func commandMap(c *config, args ...string) error {
   
	locationsResp, err := c.pokeapiClient.GetLocations(c.nextLocationsURL)	
	  if err != nil {
	     return err
	  }
    
		c.nextLocationsURL = locationsResp.Next
		c.prevLocationsURL = locationsResp.Previous 
    
		for _, location := range locationsResp.Results {
			fmt.Println(location.Name)

		}

		return nil
}




