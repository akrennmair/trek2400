package main

import (
	"math"
	"math/rand"
)

func ranf(max int) int {
	if max <= 0 {
		return 0
	}

	return rand.Intn(max)
}

func franf() float64 {
	return math.Abs(rand.NormFloat64())
}
