package main

import (
	"fmt"
	"os"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
)

func commandExit(cfg *appcfg.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}