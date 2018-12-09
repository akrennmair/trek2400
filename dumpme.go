package main

import "fmt"

func dumpme(flag bool) {
	var (
		x float64
		e *event
	)

	ship.quadx = ranf(NQUADS)
	ship.quady = ranf(NQUADS)
	ship.sectx = ranf(NSECTS)
	ship.secty = ranf(NSECTS)

	x = 1.5 * franf()
	if flag {
		fmt.Printf("%s falls into a black hole.\n", ship.shipname)
	} else {
		fmt.Printf("Computer applies full reverse power to avoid hitting the\n")
		fmt.Printf("   negative energy barrier.  A space warp was entered.\n")
	}
	/* bump repair dates forward */
	for i := 0; i < MAXEVENTS; i++ {
		e = &eventList[i]
		if e.evcode != E_FIXDV {
			continue
		}
		reschedule(e, e.date-now.date+x)
	}
	events(true)
	fmt.Printf("You are now in quadrant %d,%d.  It is stardate %.2f\n", ship.quadx, ship.quady, now.date)
	move.time = 0
}
