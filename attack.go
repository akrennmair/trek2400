package main

import (
	"fmt"
	"math"
)

func attack(resting bool) {
	var (
		hit, i, l                int
		maxhit, tothit, shldabsb int
		chgfac, propor, extradm  float64
		dustfac, tothe           float64
		cas                      int
		hitflag                  int
	)

	if move.free {
		return
	}

	if etc.enemyCount <= 0 || quad[ship.quadx][ship.quady].stars < 0 {
		return
	}

	if ship.cloaked && ship.cloakgood {
		return
	}

	enemyMove(false)
	if ship.cond == DOCKED {
		if !resting {
			fmt.Printf("Starbase shields protect the %s\n", ship.shipname)
		}
		return
	}

	chgfac = 1.0
	if move.shldchg {
		chgfac = 0.25 + 0.50*franf()
	}
	maxhit = 0.0
	tothit = 0.0
	hitflag = 0

	for i = 0; i < etc.enemyCount; i++ {
		if etc.enemyList[i].power < 20 {
			continue
		}
		if hitflag == 0 {
			fmt.Printf("\nStardate %.2f: %s attack:\n", now.date, period.enemy)
			hitflag++
		}

		/* complete the hit */
		dustfac = 0.90 + 0.01*franf()
		tothe = etc.enemyList[i].avgdist
		hit = int(float64(etc.enemyList[i].power) * math.Pow(dustfac, tothe) * param.hitfac)
		/* deplete his energy */
		dustfac = float64(etc.enemyList[i].power)
		etc.enemyList[i].power = int(dustfac * param.phasfac * (1.0 + (franf()-0.5)*0.2))
		/* see how much of hit shields will absorb */
		shldabsb = 0.0
		if ship.shldup || move.shldchg {
			propor = float64(ship.shield) / float64(param.shield)
			shldabsb = int(propor * chgfac * float64(hit))
			if shldabsb > ship.shield {
				shldabsb = ship.shield
			}
			ship.shield -= shldabsb
		}
		/* actually do the hit */
		fmt.Printf("\aHIT: %d units", int(hit))
		if !damaged(SRSCAN) {
			fmt.Printf(" from %d,%d", etc.enemyList[i].x, etc.enemyList[i].y)
		}
		cas = (shldabsb * 100) / hit
		hit -= shldabsb
		if shldabsb > 0 {
			fmt.Printf(", shields absorb %d%%, effective hit %d\n", int(cas), int(hit))
		} else {
			fmt.Printf("\n")
		}
		tothit += hit
		if hit > maxhit {
			maxhit = hit
		}
		ship.energy -= int(hit)
		if hit >= (15-game.skill)*(25-ranf(12)) {
			fmt.Printf("\aCRITICAL HIT!!!\a\n")
			/* select a device from probability vector */
			cas = ranf(1000)
			l = 0
			for ; cas >= 0; l++ {
				cas -= int(param.damprob[l])
			}
			l -= 1
			extradm = (float64(hit)*param.damfac[l])/float64(75+ranf(25)) + 0.5
			/* damage the device */
			damage(l, extradm)
			if damaged(SHIELD) {
				if ship.shldup {
					fmt.Printf("%s: Shields knocked down, captain.\n", period.helmsman)
				}
				ship.shldup = false
				move.shldchg = false
			}
		}
		if ship.energy <= 0 {
			lose(L_DSTRYD)
		}
	}

	if maxhit >= 200 || tothit >= 500 {
		cas = int(float64(tothit) * 0.015 * franf())
		if cas >= 2 {
			fmt.Printf("%s: we suffered %d casualties in that attack.\n", period.doctor, int(cas))
			game.deaths += int(cas)
			ship.crew -= int(cas)
		}
	}

	/* allow enemy to move after attacking */
	enemyMove(true)
}
