package main

import (
	"math/rand"
)

func ranf(max int) int {
	if max <= 0 {
		return 0
	}

	return rand.Intn(max)
}

func franf() float64 {
	return rand.Float64()
}
