package main

import (
	"fmt"
	"sort"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/pokeapi"
)

func commandPokedex(cfg *appcfg.Config, args ...string) error {
	if len(cfg.CaughtPokemon) == 0 {
        fmt.Println("Your Pokedex is empty.")
        return nil
    }

    list := make([]pokeapi.Pokemon, 0, len(cfg.CaughtPokemon))
    for _, p := range cfg.CaughtPokemon {
        list = append(list, p)
    }

    sort.Slice(list, func(i, j int) bool {
        return list[i].ID < list[j].ID
    })

    fmt.Println("Your Pokedex:")
    for _, p := range list {
        fmt.Printf(" - %s (id: %d)\n", p.Name, p.ID)
    }
    return nil
}