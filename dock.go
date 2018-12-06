package main

import "fmt"

func dock(v int) {
	// TODO: implement
}

func undock(_ int) {
	if ship.cond != DOCKED {
		fmt.Printf("Sulu: Pardon me captain, but we are not docked.\n")
		return
	}

	ship.cond = GREEN
	move.free = false

	/* reschedule device repair times (again) */
	for i := 0; i < MAXEVENTS; i++ {
		e := &eventList[i]
		if e.evcode != E_FIXDV {
			continue
		}
		reschedule(e, (e.date-now.date)/param.dockfac)
	}
}
