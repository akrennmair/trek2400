package main

import "fmt"

func setwarp(v int) {
	warpfac := getfltpar("Warp factor")
	if warpfac < 0.0 {
		return
	}
	if warpfac < 1.0 {
		fmt.Printf("Minimum warp speed is 1.0\n")
		return
	}
	if warpfac > 10.0 { // TODO: warp 10 should be impossible
		fmt.Printf("Maximum speed is warp 10.0\n")
		return
	}
	if warpfac > 6.0 {
		fmt.Printf("Damage to warp engines may occur above warp 6.0\n")
	}
	ship.warp = warpfac
	ship.warp2 = ship.warp * warpfac
	ship.warp3 = ship.warp2 * warpfac
}
