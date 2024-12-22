package main

import (
	"fmt"
	"errors"
)



func commandMapb(c *config, args ...string) error {
	   if c.prevLocationsURL == nil {
			 return errors.New("you are on the first page")
		 }

    locationsResp, err := c.pokeapiClient.GetLocations(c.prevLocationsURL)
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
