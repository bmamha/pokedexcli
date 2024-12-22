package main

import (
	"fmt"
	"math/rand/v2"
	"errors"
)


func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You need to enter name of pokemon")
	}
pokemonname := args[0]

fmt.Printf("Throwing a Pokeball at %v...\n", pokemonname)

pokemonResp, err := c.pokeapiClient.GetPokemon(pokemonname)
if err != nil {
	     return err
	  }

//with maximum for BaseExperience at 255. we will use that as our
// bench mark for the random variable 
experience := pokemonResp.BaseExperience

fmt.Printf("%v has %v Base Experience\n", pokemonname, experience)

catch_likelihood := rand.IntN(experience)

if catch_likelihood > 30 {

	fmt.Printf("%v escaped!\n", pokemonname)
	return nil 
} 

	fmt.Printf("%v was caught!\n", pokemonname)

	c.pokeapiClient.Pokedex[pokemonname] = pokemonResp

	return nil 

}

