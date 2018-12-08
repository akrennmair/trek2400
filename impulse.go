package main

import "fmt"

func impulse(v int) {
	if ship.cond == DOCKED {
		fmt.Printf("Scotty: Sorry captain, but we are still docked.\n")
		return
	}

	if damaged(IMPULSE) {
		out(IMPULSE)
		return
	}

	var (
		course int
		dist   float64
	)

	if getcodi(&course, &dist) {
		return
	}

	power := 20 + 100*dist
	percent := int(100*power/float64(ship.energy) + 0.5)
	if percent >= 85 {
		fmt.Printf("Scotty: That would consume %d%% of our remaining energy.\n", percent)
		if !getynpar("Are you sure that is wise") {
			return
		}
		fmt.Printf("Aye aye, sir\n")
	}
	pTime := dist / 0.095
	percent = int(100*pTime/now.time + 0.5)
	if percent >= 85 {
		fmt.Printf("Spock: That would take %d%% of our remaining time.\n", percent)
		if !getynpar("Are you sure that is wise") {
			return
		}
		fmt.Printf("(He's finally gone mad)\n")
	}
	move.time = domove(0, course, pTime, 0.095)
	ship.energy -= int(20 + 100*move.time*0.095)
}
