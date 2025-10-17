package exit

import (
	"fmt"
	"os"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
)

func init() {
	core.RegisterCommand("exit", core.Command{
		Name:        "exit                      ",
        Description: "Exit the Pokédex",
        Callback:    CommandExit,
	})
}

func CommandExit(cfg *appcfg.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}