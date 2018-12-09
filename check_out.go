package main

import "fmt"

func check_out(device int) bool {
	/* check for device ok */
	if !damaged(device) {
		return false
	}

	/* report it as being dead */
	out(device)

	/* but if we are docked, we can go ahead anyhow */
	if ship.cond != DOCKED {
		return true
	}
	fmt.Printf("  Using starbase %s\n", devices[device].name)
	return false
}
