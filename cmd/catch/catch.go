package catch

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
	"github.com/vitlobo/pokedexcli/internal/util"
)

func init() {
	core.RegisterCommand("catch", core.Command{
		Name:        "catch <pokemon_name>      ",
		Description: "Attempt to catch a Pokémon",
		Callback:    CommandCatch,
	})
}

func CommandCatch(cfg *appcfg.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	
	key := args[0]
	if _, exists := cfg.CaughtPokemon[key]; exists {
		fmt.Printf("You already caught %s.\n", util.TitleCase(key))
		return nil
	}

	p, err := cfg.PokeapiClient.GetPokemon(key)
	if err != nil { return err }

	s, err := cfg.PokeapiClient.GetPokemonSpecies(key)
	if err != nil { return err }

	// add species info
	p.CaptureRate = s.CaptureRate
	p.GenderRate = s.GenderRate
	p.GrowthRate = s.GrowthRate.Name
	p.FlavorText = getFlavor(s.FlavorTextEntries)

	// UI polish
	catchUI("Pokeball", p.Name)

	roll := rand.IntN(256)
	a := getCatchRoll(cfg, p, "pokeball")
	
	if roll > a {
		fmt.Printf("%s broke free!\n", util.TitleCase(p.Name))

		// Increase luck slightly on failure, clamp if necessary
		cfg.PlayerLuck =util.Clamp(cfg.PlayerLuck*LuckBoostOnFail, MinLuck, MaxLuck)
		return nil
	}

	fmt.Printf("%s was caught!\n", util.TitleCase(p.Name))
	fmt.Println("You may now inspect it with the 'inspect' command.")

	// Decrease luck slightly after success, clamp if necessary
	cfg.PlayerLuck = util.Clamp(cfg.PlayerLuck*LuckDecayOnSuccess, MinLuck, MaxLuck)

	cfg.CaughtPokemon[key] = p
	return nil
}