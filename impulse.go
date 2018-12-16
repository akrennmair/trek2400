package main

import "fmt"

func impulse(v int) {
	var (
		course       int
		power        int
		dist, p_time float64
		percent      int
	)

	if ship.cond == DOCKED {
		fmt.Printf("%s: Sorry captain, but we are still docked.\n", period.engineer)
		return
	}

	if damaged(IMPULSE) {
		out(IMPULSE)
		return
	}

	if getcodi(&course, &dist) {
		return
	}

	power = int(20 + 100*dist)
	percent = int(float64(100*power/ship.energy) + 0.5)
	if percent >= 85 {
		fmt.Printf("%s: That would consume %d%% of our remaining energy.\n", period.engineer, percent)
		if !getynpar("Are you sure that is wise") {
			return
		}
		fmt.Printf("Aye aye, sir\n")
	}
	p_time = dist / 0.095
	percent = int(100*p_time/now.time + 0.5)
	if percent >= 85 {
		fmt.Printf("%s: That would take %d%% of our remaining time.\n", period.firstOfficer, percent)
		if !getynpar("Are you sure that is wise") {
			return
		}
		fmt.Printf("(He's finally gone mad)\n")
	}
	move.time = domove(0, course, p_time, 0.095)
	ship.energy -= int(20 + 100*move.time*0.095)
}
