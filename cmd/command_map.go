package cmd

import (
	"errors"
	"fmt"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
)

func CommandMapf(cfg *appcfg.Config, args ...string) error {
	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil { return err }

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandMapb(cfg *appcfg.Config, args ...string) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil { return err }

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}