package main

import (
	"log"
	"time"

	_ "github.com/vitlobo/pokedexcli/cmd/catch"
	_ "github.com/vitlobo/pokedexcli/cmd/exit"
	_ "github.com/vitlobo/pokedexcli/cmd/explore"
	_ "github.com/vitlobo/pokedexcli/cmd/help"
	_ "github.com/vitlobo/pokedexcli/cmd/inspect"
	_ "github.com/vitlobo/pokedexcli/cmd/mapcmd"
	_ "github.com/vitlobo/pokedexcli/cmd/pokedex"
	_ "github.com/vitlobo/pokedexcli/cmd/save"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
	"github.com/vitlobo/pokedexcli/internal/pokeapi"
	"github.com/vitlobo/pokedexcli/internal/pokesave"
	"github.com/vitlobo/pokedexcli/internal/repl"
)

func main() {
	path, err := pokesave.DefaultPath()
	if err != nil { log.Fatal(err)}

	snap, err := pokesave.Load(path)
	if err != nil { log.Printf("load save: %v", err)}

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &appcfg.Config{
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
		PokeapiClient: pokeClient,
	}

	appcfg.ApplySnapshot(cfg, snap)

	commands := core.GetCommands()
	repl.StartRepl(cfg,commands)
}