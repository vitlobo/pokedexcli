package appcfg

import (
	"sort"
	"strings"

	"github.com/vitlobo/pokedexcli/internal/pokeapi"
	"github.com/vitlobo/pokedexcli/internal/pokesave"
)

func SnapshotFromConfig(cfg *Config) pokesave.SaveV1 {
    snap := pokesave.SaveV1{Version: pokesave.SaveVersion}
    snap.Entries = make([]pokesave.PokeEntry, 0, len(cfg.CaughtPokemon))

    names := make([]string, 0, len(cfg.CaughtPokemon))
    for name := range cfg.CaughtPokemon {
        names = append(names, name)
    }
    sort.Strings(names)

    for _, key := range names {
        p := cfg.CaughtPokemon[key]
        name := p.Name
        if name == "" {
            name = key
        }
        snap.Entries = append(snap.Entries, pokesave.PokeEntry{
            ID: p.ID, Name: name,
        })
    }
    return snap
}

func ApplySnapshot(cfg *Config, s pokesave.SaveV1) {
    m := make(map[string]pokeapi.Pokemon, len(s.Entries))
    for _, e := range s.Entries {
		key := strings.ToLower(strings.TrimSpace(e.Name))
        m[key] = pokeapi.Pokemon{ID: e.ID, Name: e.Name}
    }
    cfg.CaughtPokemon = m
}