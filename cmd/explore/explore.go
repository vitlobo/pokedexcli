package explore

import (
	"errors"
	"fmt"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
)

func init() {
	core.RegisterCommand("explore", core.Command{
		Name:        "explore <location_name>   ",
        Description: "Explore a location",
        Callback:    CommandExplore,
	})
}

func CommandExplore(cfg *appcfg.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := cfg.PokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}