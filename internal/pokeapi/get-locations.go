package pokeapi

import (
	"encoding/json"
	"net/http"
)

func(c *Client) GetLocations(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
  if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url , nil)

	if err != nil {
		return LocationArea{}, err 
	}
  
	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err 
	}

	defer res.Body.Close()

	var locations LocationArea

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locations); err != nil {
		return LocationArea{}, err  
	}
	
	return locations, nil
}
