package appcfg

import (
	"github.com/vitlobo/pokedexcli/internal/pokeapi"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	PlayerLuck       float64 // starts at 1.0, can fluctuate slightly
	NextLocationsURL *string
	PrevLocationsURL *string
	CaughtPokemon    map[string]pokeapi.Pokemon
}