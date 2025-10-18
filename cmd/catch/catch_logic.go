package catch

import (
	"math"
	"math/rand/v2"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/pokeapi"
	"github.com/vitlobo/pokedexcli/internal/util"
)

// getCatchRoll calculates the "catch rate roll" value (a) inspired by the official Pokémon formula.
// It considers HP, capture rate, Pokéball modifier, difficulty, and player luck.
func getCatchRoll(cfg *appcfg.Config, p pokeapi.Pokemon, ballType string, difficulty float64) int {
	if cfg.PlayerLuck == 0 {
		cfg.PlayerLuck = 1.0
	}

	ballMod := getBallModifier(ballType)
	hp, ok := getHPBaseStat(p)
	if !ok { hp = 50 } // fallback default

	maxHP := (2 * hp) + 100

	// Base catch factor -> increases when hp is lower or Pokémon or easier to catch
	hpRatio := float64(maxHP-hp) / float64(maxHP)
	baseCatch := float64(p.CaptureRate) * (1.0 + hpRatio*1.5)

	// Apply player luck and random variance
	luckVariance := 1.0 + (rand.Float64()*0.2 - 0.1) // +-10%
	luckFactor := util.Clamp(cfg.PlayerLuck*luckVariance, 0.5, 2.0)

	// Apply difficulty directly — higher difficulty reduces final catch rate
	a := (baseCatch * ballMod * luckFactor) / difficulty
	a += rand.Float64()*5 - 2.5 // subtle randomness for natural feel

	// Clamp to range [1, 255]
	return util.Clamp(int(math.Round(a)), 1, 255)
}

// Emulates the official 4-shake catch system from the Pokémon games.
// Returns whether the Pokémon was caught and how many shakes occurred.
func getCatchResult(a int) (int, bool) {
	if a >= 255 { return 3, true } //guaranteed capture

	// Compute shake threshold from the official Gen III+ formula:
	// b = 1048560 / sqrt(sqrt(16711680 / a))
	b := 1048560.0 / math.Sqrt(math.Sqrt(16711680.0/float64(a)))

	//Simulate up to 4 shakes (3 visual, 4th = capture)
	shakes := 0
	for i := 0; i < 4; i++ {
		roll := rand.Float64() * 65535.0
		if roll < b {
			if i < 3 {
				shakes++ // visual shake
			}
		} else {
			return shakes, false // broke free here
		}
	}
	return 3, true // all shakes passed -> caught
}

// Returns the capture modifier for different Pokéball types.
func getBallModifier(ballType string) float64 {
	switch ballType {
	case "greatball":
		return GreatBallMod
	case "ultraball":
		return UltraBallMod
	default:
		return PokeBallMod
	}
}

// Retrieves the HP base stat from a Pokémon's stat array.
func getHPBaseStat(p pokeapi.Pokemon) (int, bool) {
	for _, s := range p.Stats {
		if s.Stat.Name == "hp" {
			return s.BaseStat, true
		}
	}
	return 0, false
}