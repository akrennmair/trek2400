package main

import (
	"math"
)

func compkldist(f bool) {
	var (
		dx, dy int
		d      float64
		temp   float64
	)

	if etc.enemyCount == 0 {
		return
	}

	for i := 0; i < etc.enemyCount; i++ {
		/* compute distance to the enemy */
		dx = ship.sectx - etc.enemyList[i].x
		dy = ship.secty - etc.enemyList[i].y
		d = math.Sqrt(float64(dx*dx + dy*dy))

		if !f {
			temp = etc.enemyList[i].dist
			etc.enemyList[i].avgdist = 0.5 * (temp + d)
		} else {
			etc.enemyList[i].avgdist = d
		}
		etc.enemyList[i].dist = d
	}

	sortkl()
}

func sortkl() {
	var (
		t    enemy
		i, m int
		f    bool
	)

	m = etc.enemyCount - 1
	f = true
	for f {
		f = false
		for i = 0; i < m; i++ {
			if etc.enemyList[i].dist > etc.enemyList[i+1].dist {
				t = etc.enemyList[i]
				etc.enemyList[i] = etc.enemyList[i+1]
				etc.enemyList[i+1] = t
				f = true
			}
		}
	}
}
