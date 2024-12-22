package main

import (
	"fmt"
)

func commandHelp(c *config, args ...string) error {
	command_map := getCommands()
   fmt.Println("Welcome to the Pokedex!")
   fmt.Println("Usage")
	 for key, value := range command_map {
		fmt.Printf("%v: %v\n", key, value.description)
	}

	return nil 

}


