package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
)

func commandCatch(cfg *appcfg.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	
	key := strings.ToLower(strings.TrimSpace(args[0]))
	pokemon, err := cfg.PokeapiClient.GetPokemon(key)
	if err != nil { return err }

	if _, exists := cfg.CaughtPokemon[key]; exists {
		fmt.Printf("You already caught %s.\n", pokemon.Name)
		return nil
	}

	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.CaughtPokemon[key] = pokemon
	return nil
}