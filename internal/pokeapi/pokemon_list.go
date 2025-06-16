package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListPokemon -
func (c *Client) ListPokemon(area *string) (LocationArea, error) {
	url := baseURL + "/location-area/" + *area
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := LocationArea{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return LocationArea{}, fmt.Errorf("cache json: %v", err)
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("new request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, fmt.Errorf("request (do): %v", err)
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("io reading: %v", err)
	}
	if resp.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, dat)
	}

	pokemonResp := LocationArea{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return LocationArea{}, fmt.Errorf("json: %v (%s)", err, dat)
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
