package pokeapi

import (
	"net/http"
	"time"

  "github.com/bmamha/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cacheClient      pokecache.Cache
	Pokedex    map[string]Pokemon
}

// NewClient -
func NewClient(timeout, timeinterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			          Timeout: timeout,},
    cacheClient: pokecache.NewCache(timeinterval),
		Pokedex: make(map[string]Pokemon), 
	}
} 
