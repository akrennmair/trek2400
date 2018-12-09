package main

import "fmt"

// TODO: decide whether it's Faire Queene or Fairie Queene

func abandon(v int) {
	var (
		q *quadrant
		i int
		e *event
	)

	if ship.ship == QUEENE {
		fmt.Printf("You may not abandon yet Faire Queene\n")
		return
	}

	if ship.cond != DOCKED {
		if damaged(SHUTTLE) {
			out(SHUTTLE)
			return
		}
		fmt.Printf("Officers escape in shuttlecraft\n")
		/* decide on fate of crew */
		q = &quad[ship.quadx][ship.quady]
		if q.qsystemname == 0 || damaged(XPORTER) {
			fmt.Printf("Entire crew of %d left to die in outer space\n", ship.crew)
			game.deaths += ship.crew
		} else {
			fmt.Printf("Crew beams down to planet %s\n", systemname(q))
		}
	}
	/* see if you can be exchanged */
	if now.bases == 0 || game.captives < 20*game.skill {
		lose(L_CAPTURED)
	}
	/* re-outfit new ship */
	fmt.Printf("You are hereby put in charge of an antiquated but still\n")
	fmt.Printf("  functional ship, the Fairie Queene.\n")
	ship.ship = QUEENE
	ship.shipname = "Fairie Queene"
	param.energy = 3000
	ship.energy = 3000
	param.torped = 6
	ship.torped = 6
	param.shield = 1250
	ship.shield = 1250
	ship.shldup = false
	ship.cloaked = false
	ship.warp = 5.0
	ship.warp2 = 25.0
	ship.warp3 = 125.0
	ship.cond = GREEN
	/* clear out damages on old ship */
	for i := 0; i < MAXEVENTS; i++ {
		e = &eventList[i]
		if e.evcode != E_FIXDV {
			continue
		}
		unschedule(e)
	}
	/* get rid of some devices and redistribute probabilities */
	i = int(param.damprob[SHUTTLE] + param.damprob[CLOAK])
	param.damprob[SHUTTLE] = 0
	param.damprob[CLOAK] = 0
	for i > 0 {
		for j := 0; j < len(devices); j++ {
			if param.damprob[j] != 0 {
				param.damprob[j] += 1
				i--
				if i <= 0 {
					break
				}
			}
		}
	}
	/* pick a starbase to restart at */
	i = ranf(now.bases)
	ship.quadx = now.base[i].x
	ship.quady = now.base[i].y
	/* setup that quadrant */
	for {
		initquad(1)
		sect[ship.sectx][ship.secty] = EMPTY
		for i := 0; i < 5; i++ {
			ship.sectx = etc.starbase.x + ranf(3) - 1
			if ship.sectx < 0 || ship.sectx >= NSECTS {
				continue
			}
			ship.secty = etc.starbase.y + ranf(3) - 1
			if ship.secty < 0 || ship.secty >= NSECTS {
				continue
			}
			if sect[ship.sectx][ship.secty] == EMPTY {
				sect[ship.sectx][ship.secty] = QUEENE
				dock(0)
				compkldist(false)
				return
			}
		}
	}
}
