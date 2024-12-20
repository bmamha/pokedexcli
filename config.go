package main
import (

"github.com/bmamha/pokedexcli/internal/pokeapi"

"github.com/bmamha/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokeClientCache *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string   
}
