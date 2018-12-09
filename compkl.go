package main

import (
	"math"
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
	var (
		t    kling
		i, m int
		f    bool
	)

	m = etc.nkling - 1
	f = true
	for f {
		f = false
		for i = 0; i < m; i++ {
			if etc.klingon[i].dist > etc.klingon[i+1].dist {
				t = etc.klingon[i]
				etc.klingon[i] = etc.klingon[i+1]
				etc.klingon[i+1] = t
				f = true
			}
		}
	}
}
