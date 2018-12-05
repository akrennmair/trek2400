package main

import (
	"fmt"
	"math"
)

func attack(resting bool) {
	if move.free {
		return
	}

	if etc.nkling <= 0 || quad[ship.quadx][ship.quady].stars < 0 {
		return
	}

	if ship.cloaked && ship.cloakgood {
		return
	}

	klmove(false)
	if ship.cond == DOCKED {
		if !resting {
			fmt.Printf("Starbase shields protect the %s\n", ship.shipname)
		}
		return
	}

	chgfac := 1.0
	if move.shldchg {
		chgfac = 0.25 + 0.50*franf()
	}
	maxhit := 0.0
	tothit := 0.0
	hitflag := 0

	for i := 0; i < etc.nkling; i++ {
		if etc.klingon[i].power < 20 {
			continue
		}
		if hitflag == 0 {
			fmt.Printf("\nStardate %.2f: Klingon attack:\n", now.date)
			hitflag++
		}

		/* complete the hit */
		dustfac := 0.90 + 0.01*franf()
		tothe := etc.klingon[i].avgdist
		hit := float64(etc.klingon[i].power) * math.Pow(dustfac, tothe) * param.hitfac
		/* deplete his energy */
		etc.klingon[i].power = etc.klingon[i].power * int(param.phasfac*(1.0+(franf()-0.5)*0.2))
		/* see how much of hit shields will absorb */
		shldabsb := 0.0
		if ship.shldup || move.shldchg {
			propor := float64(ship.shield) / float64(param.shield)
			shldabsb := propor * chgfac * hit
			if int(shldabsb) > ship.shield {
				shldabsb = float64(ship.shield)
			}
			ship.shield -= int(shldabsb)
		}
		/* actually do the hit */
		fmt.Printf("\aHIT: %d units", hit)
		if !damaged(SRSCAN) {
			fmt.Printf("% from %d,%d", etc.klingon[i].x, etc.klingon[i].y)
		}
		cas := (shldabsb * 100) / hit
		hit -= shldabsb
		if shldabsb > 0 {
			fmt.Printf(", shields absorb %d%%, effective hit %d\n", cas, hit)
		} else {
			fmt.Printf("\n")
		}
		tothit += hit
		if hit > maxhit {
			maxhit = hit
		}
		ship.energy -= int(hit)
		if hit >= float64((15-game.skill)*(25-ranf(12))) {
			fmt.Printf("\aCRITICAL HIT!!!\a\n")
			/* select a device from probability vector */
			cas := ranf(1000)
			l := 0
			for ; cas >= 0; l++ {
				cas -= int(param.damprob[l])
			}
			l -= 1
			extradm := (hit*param.damfac[l])/float64(75+ranf(25)) + 0.5
			/* damage the device */
			damage(l, extradm)
			if damaged(SHIELD) {
				if ship.shldup {
					fmt.Printf("Sulu: Shields knocked down, captain.\n")
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
		cas := tothit * 0.015 * franf()
		if cas >= 2 {
			fmt.Printf("McCoy: we suffered %d casualties in that attack.\n", cas)
			game.deaths += int(cas)
			ship.crew -= int(cas)
		}
	}

	/* allow Klingons to move after attacking */
	klmove(true)
}
