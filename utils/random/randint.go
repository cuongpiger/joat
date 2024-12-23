package random

import (
	"math/rand"
	"time"
)

func RandomIntInRange(min, max int) int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
