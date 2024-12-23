package main

import (
	"fmt"
	"errors"
)



func commandExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You need to enter name of location-area")
	}
	name := args[0]
	locationResp, err := c.pokeapiClient.GetLocationName(name)
	if err != nil {
		return err
	}
  fmt.Printf("Exploring the %v...\n", name)
  fmt.Println("Found Pokemon")

	for _, encounters := range locationResp.PokemonEncounters {
      fmt.Println(encounters.Pokemon.Name)
	}


	return nil 

}
