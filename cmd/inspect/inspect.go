package inspect

import (
	"errors"
	"fmt"
	"strings"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
	"github.com/vitlobo/pokedexcli/internal/pokeapi"
)

func init() {
	core.RegisterCommand("inspect", core.Command{
		Name:        "inspect <pokemon_name>    ",
        Description: "View details about a caught Pokémon",
        Callback:    CommandInspect,
	})
}

func CommandInspect(cfg *appcfg.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name that you have caught")
	}

	name := args [0]
	p, ok := cfg.CaughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	full, err := getFullPokemon(cfg, p)
	if err != nil { return err }

	fmt.Println("Name:", full.Name)
	fmt.Println("Height:", full.Height)
	fmt.Println("Weight:", full.Weight)
	fmt.Println("Stats:")
	for _, stat := range full.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range full.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}

func getFullPokemon(cfg *appcfg.Config, p pokeapi.Pokemon) (pokeapi.Pokemon, error) {
	if p.Height != 0 { return p, nil }
	
	var full pokeapi.Pokemon
	var err error
	if p.ID != 0 {
		full, err = cfg.PokeapiClient.GetPokemon(fmt.Sprint(p.ID))
	} else {
		full, err = cfg.PokeapiClient.GetPokemon(p.Name)
	}
	if err != nil { return p, err }

	key := strings.ToLower(full.Name)
	cfg.CaughtPokemon[key] = full
	return full, nil
}