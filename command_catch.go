package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *appcfg.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	
	key := args[0]
	if _, exists := cfg.CaughtPokemon[key]; exists {
		fmt.Printf("You already caught %s.\n", key)
		return nil
	}

	p, err := cfg.PokeapiClient.GetPokemon(key)
	if err != nil { return err }
	s, err := cfg.PokeapiClient.GetPokemonSpecies(key)
	if err != nil { return err }

	p.CaptureRate = s.CaptureRate
	p.GenderRate = s.GenderRate
	p.GrowthRate = s.GrowthRate.Name
	p.FlavorText = getFlavor(s.FlavorTextEntries)

	a, ok := getCatchRoll(p)
	if !ok { return errors.New("error assesing capture rate") }
	roll := rand.IntN(256)
	
	fmt.Printf("Throwing a Pokeball at %s...\n", p.Name)
	if roll > a {
		fmt.Printf("%s escaped!\n", p.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", p.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.CaughtPokemon[key] = p
	return nil
}

func getCatchRoll(p pokeapi.Pokemon) (int, bool) {
	hp, ok := getHPBaseStat(p)
	if !ok { return 0, false }
	maxHP := (2 * hp) + 100

	a := ((3*maxHP - 2*hp) * p.CaptureRate) / (3 * maxHP)
    if a < 1 { a = 1 }
    if a > 255 { a = 255 }
    return a, true
}

func getHPBaseStat(p pokeapi.Pokemon) (int, bool) {
	for _, s := range p.Stats {
		if s.Stat.Name == "hp" {
			return s.BaseStat, true
		}
	}
	return 0, false
}

func getFlavor(entries []struct {
    FlavorText string `json:"flavor_text"`
    Language   struct{ Name string `json:"name"` } `json:"language"`
    Version    struct{ Name string `json:"name"` } `json:"version"`
}) string {
    for i := len(entries) - 1; i >= 0; i-- {
        if entries[i].Language.Name == "en" {
            s := strings.ReplaceAll(entries[i].FlavorText, "\n", " ")
            s = strings.ReplaceAll(s, "\f", " ")
            return strings.TrimSpace(s)
        }
    }
    return ""
}