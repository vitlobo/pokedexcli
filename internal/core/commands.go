package core

import "github.com/vitlobo/pokedexcli/internal/appcfg"

type Command struct {
	Name        string
	Description string
	Callback    func(*appcfg.Config,  ...string) error
}