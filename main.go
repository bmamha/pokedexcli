package main
import (
	"time"
  "github.com/bmamha/pokedexcli/internal/pokeapi"
  "github.com/bmamha/pokedexcli/internal/pokecache"

)
func main() {
	  pokeClient := pokeapi.NewClient(5 * time.Second)
		pokeCache := pokecache.NewCache(5 * time.Second)
		cfg := &config{
		  pokeClientCache: pokeCache,	
			pokeapiClient: pokeClient,
		}
     startRepl(cfg)
}



