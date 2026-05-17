package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	var url string
	// pick the URL
	if pageURL == nil {
		url = baseURL + "/location-area"
	} else {
		url = *pageURL
	}

	var mapRes RespShallowLocations

	// cache logic
	if val, ok := c.userCache.Get(url); ok {
		if err := json.Unmarshal(val, &mapRes); err != nil {
			return RespShallowLocations{}, err
		}

		return mapRes, nil
	}

	// build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// execute it
	res, err := c.cli.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	// read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// store in cache
	c.userCache.Add(url, body)

	// unmarshal json into go structs
	err = json.Unmarshal(body, &mapRes)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return mapRes, nil
}
