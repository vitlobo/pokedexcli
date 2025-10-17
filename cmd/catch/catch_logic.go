package catch

import (
	"math/rand/v2"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/pokeapi"
	"github.com/vitlobo/pokedexcli/internal/util"
)

// getCatchRoll calculates capture chance, accounting for HP, difficulty, luck, and Pokéball type.
func getCatchRoll(cfg *appcfg.Config, p pokeapi.Pokemon, ballType string) int {
	if cfg.PlayerLuck == 0 {
		cfg.PlayerLuck = 1.0
	}

	ballMod := getBallModifier(ballType)
	hp, ok := getHPBaseStat(p)
	if !ok { hp = 50 }

	maxHP := (2 * hp) + 100
	// Base formula
	a := float64(((3*maxHP - 2*hp) * p.CaptureRate) / (3 * maxHP))

	// Add difficulty scaling and luck randomness
	difficulty := 1.0 + float64(p.BaseExperience)/600.0
	luck := 0.9 + rand.Float64()*0.2 // +-10%

	a = a * ballMod * cfg.PlayerLuck * luck / difficulty

	// Clamp to range [1, 255]
	return util.ClampInt(int(a), 1, 255)
}

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

func getHPBaseStat(p pokeapi.Pokemon) (int, bool) {
	for _, s := range p.Stats {
		if s.Stat.Name == "hp" {
			return s.BaseStat, true
		}
	}
	return 0, false
}