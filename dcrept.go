package main

import "fmt"

func dcrept(v int) {
	var (
		f      bool
		x      float64
		m1, m2 float64
		e      *event
	)

	/* set up the magic factors to output the time till fixed */
	if ship.cond == DOCKED {
		m1 = 1.0 / param.dockfac
		m2 = 1.0
	} else {
		m1 = 1.0
		m2 = param.dockfac
	}
	fmt.Printf("Damage control report:\n")
	f = true

	for i := 0; i < MAXEVENTS; i++ {
		e = &eventList[i]
		if e.evcode != E_FIXDV {
			continue
		}

		/* output the title first time */
		if f {
			fmt.Printf("\t\t\t  repair times\n")
			fmt.Printf("device\t\t\tin flight  docked\n")
			f = false
		}

		/* compute time till fixed, then adjust by the magic factors */
		x = e.date - now.date
		fmt.Printf("%-24s%7.2f  %7.2f\n", devices[e.systemname].name, x*m1+0.005, x*m2+0.005)

		/* do a little consistancy checking */
	}

	if f {
		fmt.Printf("All device functional\n")
	}
}
