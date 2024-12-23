package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)


 func(c *Client) GetLocationName (name string) (LocationName, error) {
	url := baseURL + "/location-area/" + name 
  
		if val, ok := c.cacheClient.Get(url); ok {
		var cachedResponse LocationName 
		err := json.Unmarshal(val, &cachedResponse)
		if err != nil {
			return LocationName{}, err 
		}

		return cachedResponse, nil 
  }

	req, err := http.NewRequest("GET", url , nil)

	if err != nil {
		return LocationName{}, err 
	}
  
	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationName{}, err 
	}

	defer res.Body.Close()

	var location LocationName

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationName{}, nil
	}

	err = json.Unmarshal(data, &location)

	if err != nil {
		return LocationName{}, nil
	}

  c.cacheClient.Add(url, data)
	
	return location, nil

	


	
}
