package main

import "fmt"

func autover() {
	var (
		dist   float64
		course int
	)

	fmt.Printf("\007RED ALERT:  The %s is in a supernova quadrant\n", ship.shipname)
	fmt.Printf("***  Emergency override attempts to hurl %s to safety\n", ship.shipname)
	ship.warp = 6.0 + 2.0*franf()
	ship.warp2 = ship.warp * ship.warp
	ship.warp3 = ship.warp2 * ship.warp

	shldup := 0.0
	if ship.shldup {
		shldup = 1.0
	}

	dist = 0.75 + float64(ship.energy)/(ship.warp3*(shldup+1))
	if dist > 1.4142 {
		dist = 1.4242
	}
	course = ranf(360)
	etc.enemyCount = -1
	ship.cond = RED
	warp(-1, course, dist)
	attack(false)
}
