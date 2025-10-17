package mapcmd

import (
	"errors"
	"fmt"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
)

func init() {
	core.RegisterCommand("mapb", core.Command{
		Name:        "mapb                      ",
        Description: "Get the previous page of locations",
        Callback:    CommandMapb,
	})
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