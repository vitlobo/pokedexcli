package main

import (
	"fmt"
	"sort"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
)

func commandPokedex(cfg *appcfg.Config, args ...string) error {
	if len(cfg.CaughtPokemon) == 0 {
        fmt.Println("Your Pokedex is empty.")
        return nil
    }
    fmt.Println("Your Pokedex:")
    names := make([]string, 0, len(cfg.CaughtPokemon))
    for _, p := range cfg.CaughtPokemon { names = append(names, p.Name) }
    sort.Strings(names)
    for _, n := range names { fmt.Printf(" - %s\n", n) }
    return nil
}