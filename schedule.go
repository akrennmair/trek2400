package main

import (
	"fmt"
	"math"
)

/*
**  SCHEDULE AN EVENT
**
**      An event of type 'type' is scheduled for time NOW + 'offset'
**      into the first available slot.  'x', 'y', and 'z' are
**      considered the attributes for this event.
**
**      The address of the slot is returned.
 */

func schedule(typ int, offset float64, x, y, z int) *event {
	date := now.date + offset
	for i := 0; i < MAXEVENTS; i++ {
		e := &eventList[i]
		if e.evcode != 0 {
			continue
		}

		e.evcode = typ
		e.date = date
		e.x = x
		e.y = y
		e.systemname = z
		now.eventptr[typ] = e
		return e
	}

	panic(fmt.Sprintf("Cannot schedule event %d parm %d %d %d", typ, x, y, z))
}

/*
**  UNSCHEDULE AN EVENT
**
**      The event at slot 'e' is deleted.
 */

func unschedule(e *event) {
	now.eventptr[e.evcode&E_EVENT] = nil
	e.date = TOOLARGE
	e.evcode = 0
}

/*
**  Abreviated schedule routine
**
**      Parameters are the event index and a factor for the time
**      figure.
 */

func xsched(ev1, factor, x, y, z int) *event {
	when := -param.eventdly[ev1] * param.time * math.Log(franf()) / float64(factor)
	return schedule(ev1, when, x, y, z)
}
