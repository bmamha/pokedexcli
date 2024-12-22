package pokeapi
import (
	"encoding/json"
	"net/http"
	"io"
)

func (c *Client) GetPokemon(pokemonname string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonname 

	if val, ok := c.cacheClient.Get(url); ok {
		var cachedResponse Pokemon
		err := json.Unmarshal(val, &cachedResponse)
		if err != nil {
			return Pokemon{}, err 
      }
		  return cachedResponse, nil 
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
     return Pokemon{}, err 
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err 
	}

	defer res.Body.Close()

	var pokemon Pokemon 

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return Pokemon{}, nil
	}

	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return Pokemon{}, nil 
	}

	c.cacheClient.Add(url, data)

	return pokemon, nil 
}

