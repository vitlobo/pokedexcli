package mapcmd

import (
	"fmt"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
)

func init() {
	core.RegisterCommand("mapf", core.Command{
		Name:        "mapf                      ",
        Description: "Get the next page of locations",
        Callback:    CommandMapf,
	})
}

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