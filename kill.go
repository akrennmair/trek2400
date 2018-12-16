package main

import "fmt"

func killEnemy(ix, iy int) {
	fmt.Printf("   *** %s at %d,%d destroyed ***\n", period.enemy, ix, iy)

	/* remove the scoundrel */
	now.enemies -= 1
	sect[ix][iy] = EMPTY
	quad[ship.quadx][ship.quady].enemies -= 1
	quad[ship.quadx][ship.quady].scanned -= 100
	game.enemiesKilled += 1

	/* find the enemy in the enemies list */
	for i := 0; i < etc.enemyCount; i++ {
		if ix == etc.enemyList[i].x && iy == etc.enemyList[i].y {
			/* purge him from the list */
			etc.enemyCount -= 1
			for ; i < etc.enemyCount; i++ {
				etc.enemyList[i] = etc.enemyList[i+1]
			}
			break
		}
	}

	/* find out if that was the last one */
	if now.enemies <= 0 {
		win()
	}

	/* recompute the time left */
	now.time = now.resource / float64(now.enemies)
}

func killb(qx, qy int) {
	var (
		q *quadrant
		b *xy
	)

	q = &quad[qx][qy]

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
			fmt.Printf("%s: Starfleet command reports that the starbase in\n", period.comms)
			fmt.Printf("   quadrant %d,%d has been destroyed\n", qx, qy)
		} else {
			schedule(E_KATSB|E_GHOST, TOOLARGE, qx, qy, 0)
		}
	}
}

func killd(x, y, f int) {
	var (
		e *event
		q *quadrant
	)

	q = &quad[x][y]

	for i := 0; i < MAXEVENTS; i++ {
		e = &eventList[i]
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

func kills(x, y, f int) {
	var (
		q    *quadrant
		e    *event
		name string
	)

	if f != 0 {
		q = &quad[ship.quadx][ship.quady]
		sect[x][y] = EMPTY
		name = systemname(q)
		if name == "" {
			return
		}
		fmt.Printf("Inhabited starsystem %s at %d,%d destroyed\n", name, x, y)
		if f < 0 {
			game.killinhab += 1
		}
	} else {
		q = &quad[x][y]
	}
	if q.qsystemname&Q_DISTRESSED != 0 {
		/* distressed starsystem */
		e = &eventList[q.qsystemname&Q_SYSTEM]
		fmt.Printf("Distress call for %s invalidated\n", systemnameList[e.systemname])
		unschedule(e)
	}
	q.qsystemname = 0
	q.stars -= 1
}

var systemnameList = []string{
	"ERROR",
	"Talos IV",
	"Rigel III",
	"Deneb VII",
	"Canopus V",
	"Barkon IV",
	"Alpha Eridani II",
	"Omicron Ceti III",
	"Elba II",
	"Scalos",
	"Procyon IV",
	"Aldebaran III",
	"Argo",
	"Tau Ceti IV",
	"Eridani D",
	"Nimbus III",
	"Norpin V",
	"Centauri VII",
	"Kaelon II",
	"Sigma Draconis VI",
	"Betazed",
	"Gamma Tranguli VI",
	"Pyris VII",
	"Trialas IV",
	"Marcus XII",
	"Kreetassa",
	"Ardana",
	"Jouret IV",
	"Excalbia",
	"Axanar",
	"Ceti Alpha V",
	"Exo III",
}
