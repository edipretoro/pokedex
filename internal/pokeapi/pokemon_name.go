package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Pokemon(name *string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + *name
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, fmt.Errorf("cache json: %v", err)
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("new request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("request (do): %v", err)
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("io reading: %v", err)
	}
	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, dat)		
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, fmt.Errorf("json: %v (%s)", err, dat)
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
