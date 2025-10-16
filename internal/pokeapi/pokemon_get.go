package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
    url := baseURL + "/pokemon/" + pokemonName

    if val, ok := c.cache.Get(url); ok {
        var p Pokemon
        if err := json.Unmarshal(val, &p); err != nil {
            return Pokemon{}, fmt.Errorf("cache unmarshal: %w", err)
        }
        return p, nil
    }

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return Pokemon{}, fmt.Errorf("new request: %w", err)
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return Pokemon{}, fmt.Errorf("do request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
      io.Copy(io.Discard, resp.Body)
      return Pokemon{}, fmt.Errorf("pokemon %s: status %d", pokemonName, resp.StatusCode)
    }

    dat, err := io.ReadAll(resp.Body)
    if err != nil {
        return Pokemon{}, fmt.Errorf("read body: %w", err)
    }

    var p Pokemon
    if err := json.Unmarshal(dat, &p); err != nil {
        return Pokemon{}, fmt.Errorf("unmarshal: %w", err)
    }
    c.cache.Add(url, dat)
    return p, nil
}