package pokesave

import (
	"fmt"
	"os"
	"path/filepath"
)

const SaveVersion = 1

type PokeEntry struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SaveV1 struct {
	Version int         `json:"version"`
	Entries []PokeEntry `json:"entries"`
}

func DefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil { return "", fmt.Errorf("home dir: %w", err) }
	return filepath.Join(home, "Documents", "pokedexcli", "pokesave.json"), nil
}

func ensureDirFor(path string) error {
	dir := filepath.Dir(path)
    if err := os.MkdirAll(dir, 0700); err != nil {
        return fmt.Errorf("mkdir %s: %w", dir, err)
    }
    return nil
}