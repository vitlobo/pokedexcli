package catch

import (
	"errors"
	"fmt"
	"math"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
	"github.com/vitlobo/pokedexcli/internal/util"
)

func init() {
	core.RegisterCommand("catch", core.Command{
		Name:        "catch <pokémon_name>      ",
		Description: "Attempt to catch a Pokémon",
		Callback:    CommandCatch,
	})
}

func CommandCatch(cfg *appcfg.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokémon name")
	}
	
	key := args[0]

	// Prevent duplicate catches
	if _, exists := cfg.CaughtPokemon[key]; exists {
		fmt.Printf("You already caught %s.\n", util.TitleCase(key))
		return nil
	}

	// Fetch Pokémon data
	p, err := cfg.PokeapiClient.GetPokemon(key)
	if err != nil { return err }

	s, err := cfg.PokeapiClient.GetPokemonSpecies(key)
	if err != nil { return err }

	// Enrich Pokémon date -> add species info
	p.CaptureRate = s.CaptureRate
	p.GenderRate = s.GenderRate
	p.GrowthRate = s.GrowthRate.Name
	p.FlavorText = getFlavor(s.FlavorTextEntries)

	// --- Improved difficulty scaling ---
	// Compute difficulty based on both Base Experience and Capture Rate
	expFactor := float64(p.BaseExperience) / 380.0
	rateFactor := 255.0 / float64(util.Clamp(p.CaptureRate, 1, 255)) // lower capture rate -> higher difficulty

	// Weighted blend: capture rate influences ~55%, experience ~35%
	baseDiff := (rateFactor*0.55 + expFactor*0.35)

	// Nonlinear smoothing and clamping
	difficulty := 1.0 + (math.Pow(baseDiff, 0.56)-1.0)*1.15
	if p.BaseExperience < 150 { difficulty *= 0.9 } // if lower base exp, bring down difficulty some
	difficulty = util.Clamp(difficulty, 1.0, 2.9)

	catchRoll := getCatchRoll(cfg, p, "pokeball", difficulty)
	shakes, caught := getCatchResult(catchRoll)
	
	// Display the animated catch sequence
	catchUI("pokeball", key, difficulty, caught, shakes)
	
	if caught {
		// Decrease luck slightly after success, clamp if necessary
		cfg.PlayerLuck = util.Clamp(cfg.PlayerLuck*LuckDecayOnSuccess, MinLuck, MaxLuck)
		cfg.CaughtPokemon[key] = p
	} else {
		// Increase luck slightly on failure, clamp if necessary
		cfg.PlayerLuck =util.Clamp(cfg.PlayerLuck*LuckBoostOnFail, MinLuck, MaxLuck)
	}

	return nil
}