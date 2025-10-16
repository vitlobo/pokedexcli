package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//GetLocation
func (c *Client) GetLocation(locationName string) (Location, error) {
    url := baseURL + "/location-area/" + locationName

    if val, ok := c.cache.Get(url); ok {
        var loc Location
        if err := json.Unmarshal(val, &loc); err != nil {
            return Location{}, fmt.Errorf("cache unmarshal: %w", err)
        }
        return loc, nil
    }

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return Location{}, fmt.Errorf("new request: %w", err)
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return Location{}, fmt.Errorf("do request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        io.Copy(io.Discard, resp.Body)
        return Location{}, fmt.Errorf("location %s: status %d", locationName, resp.StatusCode)
    }

    dat, err := io.ReadAll(resp.Body)
    if err != nil {
        return Location{}, fmt.Errorf("read body: %w", err)
    }

    var loc Location
    if err := json.Unmarshal(dat, &loc); err != nil {
        return Location{}, fmt.Errorf("unmarshal: %w", err)
    }

    c.cache.Add(url, dat)
    return loc, nil
}