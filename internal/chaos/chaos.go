package chaos

import (
	"math/rand"
	"time"
)

type DefaultChaos struct{}

const chaosChance = 0.3

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (d *DefaultChaos) Hit() bool {
	return rand.Float64() < chaosChance
}
