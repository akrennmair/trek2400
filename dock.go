package main

import "fmt"

func dock(v int) {
	var (
		ok bool
		e  *event
	)

	if ship.cond == DOCKED {
		fmt.Printf("%s: But captain, we are already docked\n", names.navigator)
		return
	}

	/* check for ok to dock, i.e., adjacent to a starbase */
	ok = false
	for i := ship.sectx - 1; i <= ship.sectx+1 && !ok; i++ {
		if i < 0 || i >= NSECTS {
			continue
		}
		for j := ship.secty - 1; j <= ship.secty+1; j++ {
			if j < 0 || j >= NSECTS {
				continue
			}
			if sect[i][j] == BASE {
				ok = true
				break
			}
		}
	}
	if !ok {
		fmt.Printf("%s: But captain, we are not adjacent to a starbase.\n", names.navigator)
		return
	}

	/* restore resources */
	ship.energy = param.energy
	ship.torped = param.torped
	ship.shield = param.shield
	ship.crew = param.crew
	game.captives += param.brigfree - ship.brigfree
	ship.brigfree = param.brigfree

	/* reset ship's defenses */
	ship.shldup = false
	ship.cloaked = false
	ship.cond = DOCKED
	ship.reserves = param.reserves

	/* recalibrate space inertial navigation system */
	ship.sinsbad = false

	/* output any saved radio messages */
	dumpssradio()

	/* reschedule any device repairs */
	for i := 0; i < MAXEVENTS; i++ {
		e = &eventList[i]
		if e.evcode != E_FIXDV {
			continue
		}
		reschedule(e, (e.date-now.date)*param.dockfac)
	}
}

func undock(_ int) {
	var e *event

	if ship.cond != DOCKED {
		fmt.Printf("%s: Pardon me captain, but we are not docked.\n", names.helmsman)
		return
	}

	ship.cond = GREEN
	move.free = false

	/* reschedule device repair times (again) */
	for i := 0; i < MAXEVENTS; i++ {
		e = &eventList[i]
		if e.evcode != E_FIXDV {
			continue
		}
		reschedule(e, (e.date-now.date)/param.dockfac)
	}
}
