package main

import "fmt"

func damage(dev int, dam float64) {
	if dam <= 0.0 {
		return
	}

	fmt.Printf("\t%s damaged\n", devices[dev].name)

	/* find actual length till it will be fixed */
	if ship.cond == DOCKED {
		dam *= param.dockfac
	}
	f := damaged(dev)
	if !f {
		/* new damages -- schedule a fix */
		schedule(E_FIXDV, dam, 0, 0, dev)
		return
	}
	/* device already damaged -- add to existing damages */
	/* scan for old damages */
	for i := 0; i < MAXEVENTS; i++ {
		e := &eventList[i]
		if e.evcode != E_FIXDV || e.systemname != dev {
			continue
		}
		reschedule(e, e.date-now.date+dam)
		return
	}
	panic(fmt.Sprintf("Cannot find old damages %d\n", dev))
}
