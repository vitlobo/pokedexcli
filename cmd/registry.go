package cmd

import (
	"github.com/vitlobo/pokedexcli/internal/core"
)

func GetCommands() map[string]core.Command {
	return map[string]core.Command{
		"catch": {
			Name:        "catch <pokemon_name>      ",
			Description: "Attempt to catch a pokemon",
			Callback:    CommandCatch,
		},
		"exit": {
			Name:        "exit                      ",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"explore": {
			Name:        "explore <location_name>   ",
			Description: "Explore a location",
			Callback:    CommandExplore,
		},
		"help": {
			Name:        "help                      ",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>    ",
			Description: "View details about a caught Pokemon",
			Callback:    CommandInspect,
		},
		"map": {
			Name:        "map                       ",
			Description: "Get the next page of locations",
			Callback:    CommandMapf,
		},
		"mapb": {
			Name:        "mapb                      ",
			Description: "Get the previous page of locations",
			Callback:    CommandMapb,
		},
		"pokedex": {
			Name:        "pokedex                   ",
			Description: "See all the Pokemon you've caught",
			Callback:    CommandPokedex,
		},
		"save": {
			Name:        "save                      ",
			Description: "Save pokedex progress",
			Callback:    CommandSave,
		},
	}
}