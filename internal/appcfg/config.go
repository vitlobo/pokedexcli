package appcfg

import (
	"github.com/vitlobo/pokedexcli/internal/pokeapi"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	CaughtPokemon    map[string]pokeapi.Pokemon
}