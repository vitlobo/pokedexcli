package core

import "github.com/vitlobo/pokedexcli/internal/appcfg"

type Command struct {
	Name        string
	Description string
	Callback    func(*appcfg.Config, ...string) error
}

var commandRegistry = make(map[string]Command)

// RegisterCommand lets individual packages add themselves to the registry.
func RegisterCommand(name string, cmd Command) {
	commandRegistry[name] = cmd
}

// GetCommands returns the registered commands.
func GetCommands() map[string]Command {
	return commandRegistry
}