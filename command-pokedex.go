package main

import (
	"fmt"
)


func commandPokedex(c *config, args ...string) error {
  
	fmt.Println("Your Pokedex")

	for key, _ := range c.pokeapiClient.Pokedex {
		fmt.Printf("- %v\n", key)
	}

	return nil
}
