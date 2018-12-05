package main

import "fmt"

func check_out(device int) bool {
	if !damaged(device) {
		return false
	}

	out(device)
	if ship.cond != DOCKED {
		return true
	}
	fmt.Printf("  Using starbase %s\n", devices[device].name)
	return false
}
