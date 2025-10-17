package cmd

import (
	"fmt"

	"github.com/vitlobo/pokedexcli/internal/appcfg"
	"github.com/vitlobo/pokedexcli/internal/pokesave"
)

func CommandSave(cfg *appcfg.Config, args ...string) error {
	path, err := pokesave.DefaultPath()
    if err != nil { return err }

    snap := appcfg.SnapshotFromConfig(cfg)
    if err := pokesave.Write(path, snap); err != nil {
        return fmt.Errorf("save failed: %w", err)
    }
    fmt.Println("Data saved successfully!")
    return nil
}