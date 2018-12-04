package main

/*  DAMAGED -- check for device damaged
**
**      This is a boolean function which returns non-zero if the
**      specified device is broken.  It does this by checking the
**      event list for a "device fix" action on that device.
 */

func damaged(dev int) bool {
	d := dev

	for i := 0; i < MAXEVENTS; i++ {
		e := &events[i]
		if e.evcode != E_FIXDV {
			continue
		}
		if e.systemname == d {
			return true
		}
	}

	/* device fix not in event list -- device must not be broken */
	return false
}
