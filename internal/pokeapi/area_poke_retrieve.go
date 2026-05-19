package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(area string) (AreaEncounterInfo, error) {
	var pokeList AreaEncounterInfo

	url := baseURL + "/location-area/" + area + "/"

	// cache logic
	if val, ok := c.userCache.Get(url); ok {
		if err := json.Unmarshal(val, &pokeList); err != nil {
			return AreaEncounterInfo{}, err
		}
		return pokeList, nil
	}

	// build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaEncounterInfo{}, err
	}

	// execute it
	res, err := c.cli.Do(req)
	if err != nil {
		return AreaEncounterInfo{}, err
	}
	defer res.Body.Close()

	// read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaEncounterInfo{}, err
	}

	// store in cache
	c.userCache.Add(url, body)

	// unmarshal json into go structs
	err = json.Unmarshal(body, &pokeList)
	if err != nil {
		return AreaEncounterInfo{}, err
	}

	return pokeList, nil
}
