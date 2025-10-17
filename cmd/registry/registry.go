package registry

import (
	_ "github.com/vitlobo/pokedexcli/cmd/catch"
	_ "github.com/vitlobo/pokedexcli/cmd/exit"
	_ "github.com/vitlobo/pokedexcli/cmd/explore"
	_ "github.com/vitlobo/pokedexcli/cmd/help"
	_ "github.com/vitlobo/pokedexcli/cmd/inspect"
	_ "github.com/vitlobo/pokedexcli/cmd/mapcmd"
	_ "github.com/vitlobo/pokedexcli/cmd/pokedex"
	_ "github.com/vitlobo/pokedexcli/cmd/save"

	"github.com/vitlobo/pokedexcli/internal/core"
)

func GetCommands() map[string]core.Command {
	return core.GetCommands()
}