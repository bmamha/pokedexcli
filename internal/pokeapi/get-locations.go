package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)

func(c *Client) GetLocations(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
  if pageURL != nil {
		url = *pageURL
	}
  
	if val, ok := c.cacheClient.Get(url); ok {
		var cachedResponse LocationArea
		err := json.Unmarshal(val, &cachedResponse)
		if err != nil {
			return LocationArea{}, err 
		}

		return cachedResponse, nil 

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

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationArea{}, nil
	}

	err = json.Unmarshal(data, &locations)

	if err != nil {
		return LocationArea{}, nil
	}

  c.cacheClient.Add(url, data)
	
	return locations, nil
}  
