package catch

import (
	"fmt"
	"time"

	"github.com/vitlobo/pokedexcli/internal/util"
)

// catchUI simulates the Pokéball shaking animation in the CLI.
func catchUI(ballName, pokemonName string) {
	fmt.Printf("Throwing a %s at %s...\n", ballName, util.TitleCase(pokemonName))
	time.Sleep(1 * time.Second)

	for i := 0; i < BallShakeCount; i++ {
		fmt.Print("The ball shakes")
		for j := 0; j <= i; j++ {
			fmt.Print(".")
		}
		fmt.Println()
		time.Sleep(BallShakeDelay)
	}
}
