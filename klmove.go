package main

import "fmt"

func enemyMove(fl bool) {
	var (
		n              int
		k              *enemy
		dx, dy         float64
		nextx, nexty   int
		lookx, looky   int
		motion         int
		fudgex, fudgey int
		qx, qy         int
		bigger         float64
		i              int
	)

	for n = 0; n < etc.enemyCount; n++ {
		k = &etc.enemyList[n]
		i = 100
		if fl {
			i = 100 * (k.power / param.enemyPower)
		}

		flint := 0
		if fl {
			flint = 1
		}
		if float64(ranf(i)) >= param.moveprob[2*move.newquad+flint] {
			continue
		}

		/* compute distance to move */
		motion = ranf(75) - 25
		motion = int(float64(motion) * (k.avgdist * param.movefac[2*move.newquad+flint]))
		/* compute direction */
		dx = float64(ship.sectx - k.x + ranf(3) - 1)
		dy = float64(ship.secty - k.y + ranf(3) - 1)
		bigger = dx
		if dy > bigger {
			bigger = dy
		}
		if bigger == 0.0 {
			bigger = 1.0
		}
		dx = float64(dx)/float64(bigger) + 0.5
		dy = float64(dy)/float64(bigger) + 0.5
		if motion < 0 {
			motion = -motion
			dx = -dx
			dy = -dy
		}
		fudgex, fudgey = 1, 1
		/* try to move the enemy */
		nextx = k.x
		nexty = k.y
		for ; motion > 0; motion-- {
			lookx = int(float64(nextx) + dx)
			looky = int(float64(nexty) + dy)
			if lookx < 0 || lookx >= NSECTS || looky < 0 || looky >= NSECTS {
				/* new quadrant */
				qx = ship.quadx
				qy = ship.quady
				if lookx < 0 {
					qx -= 1
				} else if lookx >= NSECTS {
					qx += 1
				}

				if looky < 0 {
					qy -= 1
				} else if looky >= NSECTS {
					qy += 1
				}

				if qx < 0 || qx >= NQUADS || qy < 0 || qy >= NQUADS || quad[qx][qy].stars < 0 || quad[qx][qy].enemies > MAXKLQUAD-1 {
					break
				}

				if !damaged(SRSCAN) {
					fmt.Printf("%s at %d,%d escapes to quadrant %d,%d\n",
						names.enemy, k.x, k.y, qx, qy)
					motion = quad[qx][qy].scanned
					if motion >= 0 && motion < 1000 {
						quad[qx][qy].scanned += 100
					}
					motion = quad[ship.quadx][ship.quady].scanned
					if motion >= 0 && motion < 1000 {
						quad[ship.quadx][ship.quady].scanned -= 100
					}
				}
				sect[k.x][k.y] = EMPTY
				quad[qx][qy].enemies += 1
				etc.enemyCount -= 1
				*k = etc.enemyList[etc.enemyCount]
				quad[ship.quadx][ship.quady].enemies -= 1
				k = nil
				break
			}
			if sect[lookx][looky] != EMPTY {
				lookx = nextx + fudgex
				if lookx < 0 || lookx >= NSECTS {
					lookx = int(float64(nextx) + dx)
				}
				if sect[lookx][looky] != EMPTY {
					fudgex = -fudgex
					looky = nexty + fudgey
					if looky < 0 || looky >= NSECTS || sect[lookx][looky] != EMPTY {
						fudgey = -fudgey
						break
					}
				}
			}
			nextx = lookx
			nexty = looky
		}
		if k != nil && (k.x != nextx || k.y != nexty) {
			if !damaged(SRSCAN) {
				fmt.Printf("%s at %d,%d moves to %d,%d\n", names.enemy, k.x, k.y, nextx, nexty)
			}
			sect[k.x][k.y] = EMPTY
			k.x, k.y = nextx, nexty
			sect[nextx][nexty] = ENEMY
		}
	}
	compkldist(false)
}
