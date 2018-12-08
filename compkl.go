package main

import (
	"math"
	"sort"
)

func compkldist(f bool) {
	if etc.nkling == 0 {
		return
	}

	var temp float64

	for i := 0; i < etc.nkling; i++ {
		dx := ship.sectx - etc.klingon[i].x
		dy := ship.secty - etc.klingon[i].y
		d := math.Sqrt(float64(dx*dx + dy*dy))

		if !f {
			temp = etc.klingon[i].dist
			etc.klingon[i].avgdist = 0.5 * (temp + d)
		} else {
			etc.klingon[i].avgdist = d
		}
		etc.klingon[i].dist = d
	}

	sortkl()
}

func sortkl() {
	sort.Slice(etc.klingon[:], func(i, j int) bool {
		return etc.klingon[i].dist < etc.klingon[j].dist
	})
}
