package main

import (
	"fmt"
	"errors"
)

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You need to enter name of pokemon")
	}

	pokemonName := args[0]

	pokemon, caught := c.pokeapiClient.Pokedex[pokemonName]

	if caught {
		fmt.Printf("Name: %v\n", pokemonName)	
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")

		for _, s := range pokemon.Stats {
			fmt.Printf("-%v: %v\n",s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types")

		for _, t := range pokemon.Types {
			fmt.Printf("- %v\n", t.Type.Name)
		}
	} else {
		fmt.Println("You have not caught the pokemon!")
	}

	return nil
}
