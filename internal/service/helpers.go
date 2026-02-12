package service

import "math/rand"

func randBool() bool {
	return rand.Intn(2) == 1
}
