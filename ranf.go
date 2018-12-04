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
	// TODO: find out what the original function does.
	//double          t;
	//t = random() & 077777;
	//return (t / 32767.0);
	return math.Abs(rand.NormFloat64())
}
