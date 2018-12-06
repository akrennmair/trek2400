package main

import "fmt"

func killk(ix, iy int) {
	fmt.Printf("   *** Klingon at %d,%d destroyed ***\n", ix, iy)

	/* remove the scoundrel */
	now.klings -= 1
	sect[ix][iy] = EMPTY
	quad[ship.quadx][ship.quady].klings -= 1
	quad[ship.quadx][ship.quady].scanned -= 100
	game.killk += 1

	/* find the Klingon in the Klingon list */
	for i := 0; i < etc.nkling; i++ {
		if ix == etc.klingon[i].x && iy == etc.klingon[i].y {
			/* purge him from the list */
			etc.nkling -= 1
			for ; i < etc.nkling; i++ {
				etc.klingon[i] = etc.klingon[i+1]
			}
			break
		}
	}

	/* find out if that was the last one */
	if now.klings <= 0 {
		win()
	}

	/* recompute the time left */
	now.time = now.resource / float64(now.klings)
}

func killb(qx, qy int) {
	q := &quad[qx][qy]

	if q.bases <= 0 {
		return
	}

	if !damaged(SSRADIO) {
		/* then update starchart */
		if q.scanned < 1000 {
			q.scanned -= 10
		} else if q.scanned > 1000 {
			q.scanned = -1
		}
	}

	q.bases = 0
	now.bases -= 1
	var b *xy
	for i := 0; i < MAXBASES; i++ {
		b = &now.base[i]
		if qx == b.x && qy == b.y {
			break
		}
	}
	*b = now.base[now.bases]
	if qx == ship.quadx && qx == ship.quady {
		sect[etc.starbase.x][etc.starbase.y] = EMPTY
		if ship.cond == DOCKED {
			undock(0)
		}
		fmt.Printf("Starbase at %d,%d destroyed\n", etc.starbase.x, etc.starbase.y)
	} else {
		if !damaged(SSRADIO) {
			fmt.Printf("Uhura: Starfleet command reports that the starbase in\n")
			fmt.Printf("   quadrant %d,%d has been destroyed\n", qx, qy)
		} else {
			schedule(E_KATSB|E_GHOST, TOOLARGE, qx, qy, 0)
		}
	}
}

func killd(x, y, f int) {
	q := &quad[x][y]

	for i := 0; i < MAXEVENTS; i++ {
		e := &eventList[i]
		if e.x != x || e.y != y {
			continue
		}
		switch e.evcode {
		case E_KDESB:
			if f != 0 {
				fmt.Printf("Distress call for starbase in %d,%d nullified\n", x, y)
				unschedule(e)
			}
			break
		case E_ENSLV, E_REPRO:
			if f != 0 {
				fmt.Printf("Distress call for %s in quadrant %d,%d nullified\n",
					systemnameList[e.systemname], x, y)
				q.qsystemname = e.systemname
				unschedule(e)
			} else {
				e.evcode |= E_GHOST
			}
		}
	}
}

var systemnameList = []string{
	"ERROR",
	"Talos IV",
	"Rigel III",
	"Deneb VII",
	"Canopus V",
	"Icarus I",
	"Prometheus II",
	"Omega VII",
	"Elysium I",
	"Scalos IV",
	"Procyon IV",
	"Arachnid I",
	"Argo VIII",
	"Triad III",
	"Echo IV",
	"Nimrod III",
	"Nemisis IV",
	"Centarurus I",
	"Kronos III",
	"Spectros V",
	"Beta III",
	"Gamma Tranguli VI",
	"Pyris III",
	"Triachus",
	"Marcus XII",
	"Kaland",
	"Ardana",
	"Stratos",
	"Eden",
	"Arrikis",
	"Epsilon Eridani IV",
	"Exo III",
}
