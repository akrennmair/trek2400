package main

import "fmt"

func dumpssradio() int {
	chkrest := 0
	for j := 0; j < MAXEVENTS; j++ {
		e := &eventList[j]
		/* if it is not hidden, then just ignore it */
		if e.evcode&E_HIDDEN == 0 {
			continue
		}
		if e.evcode&E_GHOST != 0 {
			unschedule(e)
			fmt.Printf("Starsystem %s in quadrant %d,%d is no longer distressed\n",
				systemname(&quad[e.x][e.y]), e.x, e.y)
		}

		switch e.evcode {
		case E_KDESB:
			fmt.Printf("Starbase in quadrant %d,%d is under attack\n", e.x, e.y)
			chkrest++
		case E_ENSLV, E_REPRO:
			fmt.Printf("Starsystem %s in quadrant %d,%d is distressed\n",
				systemname(&quad[e.x][e.y]), e.x, e.y)
			chkrest++
		}
	}

	return chkrest
}
