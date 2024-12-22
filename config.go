package main
import (

"github.com/bmamha/pokedexcli/internal/pokeapi"

)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string   
}
