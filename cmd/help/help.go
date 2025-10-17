package help

import (
	"fmt"
	"sort"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/core"
)

func init() {
	core.RegisterCommand("help", core.Command{
		Name:        "help                      ",
		Description: "Displays this help message",
		Callback:    CommandHelp,
	})
}

func CommandHelp(cfg *appcfg.Config, args ...string) error {
	cmds := core.GetCommands()

	keys := make([]string, 0, len(cmds))
	for k := range cmds { keys = append(keys, k) }
	sort.Strings(keys)

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, k := range keys {
		c := cmds[k]
		fmt.Printf(" - %-16s %s\n", c.Name, c.Description)
	}
	fmt.Println()
	return nil
}