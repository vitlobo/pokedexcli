package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonSpecies(pokemonName string) (PokemonSpecies, error) {
	url := baseURL + "/pokemon-species/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		var s PokemonSpecies
        if err := json.Unmarshal(val, &s); err != nil {
            return PokemonSpecies{}, fmt.Errorf("cache unmarshal: %w", err)
        }
        return s, nil
    }

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return PokemonSpecies{}, fmt.Errorf("new request: %w", err)
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonSpecies{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		io.Copy(io.Discard, resp.Body)
		return PokemonSpecies{}, fmt.Errorf("species %s: status %d", pokemonName, resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonSpecies{}, fmt.Errorf("read body: %w", err)
	}

	var s PokemonSpecies
	if err := json.Unmarshal(dat, &s); err != nil {
		return PokemonSpecies{}, fmt.Errorf("unmarshal: %w", err)
	}
	c.cache.Add(url, dat)
	return s, nil
}