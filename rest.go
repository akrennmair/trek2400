package main

import "fmt"

func rest(v int) {
	var (
		t       float64
		percent int
	)

	t = getfltpar("How long")
	if t <= 0.0 {
		return
	}

	percent = int(100*t/now.time + 0.5)
	if percent >= 70 {
		fmt.Printf("%s: That would take %d%% of our remaining time.\n", names.firstOfficer, percent)
		if !getynpar("Are you really certain that is wise") {
			return
		}
	}
	move.time = t

	/* boundary condition is the LRTB */
	t = now.eventptr[E_LRTB].date - now.date
	if ship.cond != DOCKED && move.time > t {
		move.time = t + 0.0001
	}
	move.free = false
	move.resting = true
}
