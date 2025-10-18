package catch

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/vitlobo/pokedexcli/internal/util"
)

// catchUI simulates a Pokéball throw animation
// with pacing, shake variation, and suspense based on difficulty and actual shake count.
func catchUI(ballType, pokemonName string, difficulty float64, caught bool, shakes int) {
	displayName := util.TitleCase(pokemonName)

	announceThrow(ballType, displayName, difficulty)
	delay := computeDelay(ballType, difficulty)
	performShakeSequence(delay, shakes)
	showOutcome(displayName, delay, caught, shakes)
}

// Displays the throw message and Pokéball flair animation
func announceThrow(ballType, displayName string, difficulty float64) {
	fmt.Printf("You throw a %s at %s!\n", util.TitleCase(ballType), displayName)
	time.Sleep(1500 * time.Millisecond)

	fmt.Println(getThrowMessage(ballType, difficulty))
	time.Sleep(1500 * time.Millisecond)
}

// Calculates pacing based on difficulty and Pokéball modifiers
func computeDelay(ballType string, difficulty float64) time.Duration {
	// Clamp difficulty range for smoother scaling (1.0-3.0)
	difficulty = util.Clamp(difficulty, 1.0, 3.0)

	// Base delay scales with difficulty (harder Pokémon = slower shakes)
	delay := time.Duration(float64(1600*time.Millisecond) * (1.0 + 0.1*(difficulty-1.0)))
	switch ballType {
	case "greatball":
		delay = time.Duration(float64(delay) * 0.9)
	case "ultraball":
		delay = time.Duration(float64(delay) * 0.8)
	}
	return delay
}

func getThrowMessage(ballType string, difficulty float64) string {
	var ball string
	switch ballType {
	case "greatball":
		ball = "Great Ball"
	case "ultraball":
		ball = "Ultra Ball"
	default:
		ball = "Poké Ball"
	}

	switch {
	case difficulty < 1.3:
		// Easy/common Pokémon
		easy := []string{
			fmt.Sprintf("You toss the %s lightly — it arcs smoothly through the air.", ball),
			fmt.Sprintf("The %s sails effortlessly toward the Pokémon.", ball),
			fmt.Sprintf("A clean throw! The %s glides perfectly onto target.", ball),
		}
		return easy[rand.IntN(len(easy))]

	case difficulty < 2.2:
		// Mid-tier difficulty
		medium := []string{
			fmt.Sprintf("You hurl the %s with focus — it spins toward the Pokémon.", ball),
			fmt.Sprintf("The %s zips forward with a satisfying spin.", ball),
			fmt.Sprintf("Nice throw! The %s flashes as it closes in on the target.", ball),
		}
		return medium[rand.IntN(len(medium))]

	default:
		// Hard or legendary-tier
		hard := []string{
			fmt.Sprintf("You launch the %s with all your strength — it cuts through the air!", ball),
			fmt.Sprintf("The %s rockets forward, trembling from the force of the throw.", ball),
			fmt.Sprintf("You grip the %s tightly and throw hard — this one has to count!", ball),
		}
		return hard[rand.IntN(len(hard))]
	}
}

// Runs the shake animation and returns how many times it shook
func performShakeSequence(delay time.Duration, shakes int) {
	labels := [][]string{
		{"once.", "gently once."},
		{"twice..", "twice, wobbling slightly.."},
		{"thrice...", "a third time...!"},
	}

	for i := 0; i < shakes && i < 3; i++ {
		//fmt.Printf("The ball shakes %s\n", labels[i])
		fmt.Printf("The ball shakes %s\n", labels[i][rand.IntN(len(labels[i]))])
		time.Sleep(delay)
	}
}

// Shows a randomized failure flavor message.
func printFailMessage(shakes int) {
	switch shakes {
	case 0:
		fmt.Println("Oh no! The Pokémon broke free immediately!")
	case 1:
		fmt.Println("Aww! It appeared to be caught!")
	case 2:
		fmt.Println("Aargh! Almost had it!")
	case 3:
		fmt.Println("Shoot! It was so close too!")
	default:
		fmt.Println("The Pokémon broke free!")
	}
}

// Handles suspense and prints success or failure message
func showOutcome(displayName string, delay time.Duration, caught bool, shakes int) {
	// Suspense before outcome if shakes >= 3
	if shakes >= 3 { time.Sleep(delay / 3) }

	if caught {
		fmt.Println()
		fmt.Println("*click*")
		fmt.Println()
		time.Sleep(700 * time.Millisecond)
		fmt.Printf("Gotcha! %s was caught!\n", displayName)
		fmt.Printf("You may now view it with the 'inspect %s' command.\n", displayName)
	} else {
		printFailMessage(shakes)
	}
}