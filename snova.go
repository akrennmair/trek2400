package main

import (
	"fmt"
	"time"
)

func snova(x, y int) {
	var (
		qx, qy int
		ix, iy int
		f      int
		dx, dy int
		n      int
		q      *quadrant
	)

	ix = x
	if ix < 0 {
		/* choose a quadrant */
		for {
			qx = ranf(NQUADS)
			qy = ranf(NQUADS)
			q = &quad[qx][qy]
			if q.stars > 0 {
				break
			}
		}
		if ship.quadx == qx && ship.quady == qy {
			/* select a particular star */
			n = ranf(q.stars)
			for ix = 0; ix < NSECTS; ix++ {
				for iy = 0; iy < NSECTS; iy++ {
					if sect[ix][iy] == STAR || sect[ix][iy] == INHABIT {
						n -= 1
						if n <= 0 {
							break
						}
					}
				}
			}
			f = 1
		}
	} else {
		iy = y
		qx = ship.quadx
		qy = ship.quady
		q = &quad[qx][qy]
		f = 1
	}

	if f != 0 {
		fmt.Printf("\a\nRED ALERT: supernova occurring at %d,%d\n", ix, iy)
		dx = ix - ship.sectx
		dy = iy - ship.secty
		if dx*dx+dy*dy <= 2 {
			fmt.Printf("***  Emergency override attem")
			time.Sleep(1 * time.Second)
			fmt.Printf("\n")
			lose(L_SNOVA)
		}
		q.scanned = 1000
	} else {
		if !damaged(SSRADIO) {
			q.scanned = 1000
			fmt.Printf("\nUhura: Captain, Starfleet Command reports a supernova\n")
			fmt.Printf("  in quadrant %d,%d.  Caution is advised\n", qx, qy)
		}
	}

	dx = q.klings
	dy = q.stars
	now.klings -= dx
	if x >= 0 {
		game.kills += dy
		if q.bases != 0 {
			killb(qx, qy)
		}
		game.killk += dx
	} else {
		if q.bases != 0 {
			killb(qx, qy)
		}
	}
	b := 0
	if x >= 0 {
		b = 1
	}
	killd(qx, qy, b)
	q.stars = -1
	q.klings = 0
	if now.klings <= 0 {
		fmt.Printf("Lucky devil, that supernova destroyed the last klingon\n")
		win()
	}
}
