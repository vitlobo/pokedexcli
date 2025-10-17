package catch

import "time"

//
// ====== Constants ======
//

// Luck modifiers
const (
	LuckDecayOnSuccess   float64 = 0.9
	LuckBoostOnFail      float64 = 1.02
	MinLuck 		     float64 = 0.75
	MaxLuck              float64 = 1.25
	BallShakeCount           int = 3
	BallShakeDelay               = 600 * time.Millisecond
)

//Pokeball constants
const (
	PokeBallMod  float64 = 1.0
	GreatBallMod float64 = 1.5
	UltraBallMod float64 = 2.0
)