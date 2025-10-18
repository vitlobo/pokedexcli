package pokedex

import (
	"fmt"
	"sort"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
	"github.com/vitlobo/pokedexcli/internal/pokeapi"
)

func init() {
	core.RegisterCommand("pokedex", core.Command{
		Name:        "pokedex                   ",
        Description: "See all the Pokémon you've caught",
        Callback:    CommandPokedex,
	})
}

func CommandPokedex(cfg *appcfg.Config, args ...string) error {
	if len(cfg.CaughtPokemon) == 0 {
        fmt.Println("Your Pokédex is empty.")
        return nil
    }

    list := make([]pokeapi.Pokemon, 0, len(cfg.CaughtPokemon))
    for _, p := range cfg.CaughtPokemon {
        list = append(list, p)
    }

    sort.Slice(list, func(i, j int) bool {
        return list[i].ID < list[j].ID
    })

    fmt.Println("Your Pokédex:")
    for _, p := range list {
        fmt.Printf(" - %s (id: %d)\n", p.Name, p.ID)
    }
    return nil
}