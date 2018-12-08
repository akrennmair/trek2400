package main

import (
	"fmt"
	"math"
)

func domove(ramflag int, course int, p_time float64, speed float64) float64 {
	var (
		angle        float64
		x, y, dx, dy float64
		ix, iy       int
		bigger       float64
		n            int
		dist         float64
		sectsize     float64
		xn           float64
		evtime       float64
	)

	sectsize = NSECTS
	/* initialize delta factors for move */
	angle = float64(course) * 0.0174532925
	if damaged(SINS) {
		angle += param.navigcrud[1] * (franf() - 0.5)
	} else if ship.sinsbad {
		angle += param.navigcrud[0] * (franf() - 0.5)
	}
	dx = -math.Cos(angle)
	dy = math.Sin(angle)
	bigger = math.Abs(dx)
	dist = math.Abs(dy)
	if dist > bigger {
		bigger = dist
	}
	dx /= bigger
	dy /= bigger

	/* check for long range tractor beams */
	/****  TEMPORARY CODE == DEBUGGING  ****/
	evtime = now.eventptr[E_LRTB].date - now.date

	if p_time > evtime && etc.nkling < 3 {
		/* then we got a LRTB */
		evtime += 0.005
		p_time = evtime
	} else {
		evtime = -1.0e50
	}
	dist = p_time * speed

	/* move within quadrant */
	sect[ship.sectx][ship.secty] = EMPTY
	x = float64(ship.sectx) + 0.5
	y = float64(ship.secty) + 0.5
	xn = NSECTS * dist * bigger
	n = int(xn + 0.5)

	move.free = false

	for i := 0; i < n; i++ {
		x += dx
		y += dy
		ix := int(x)
		iy := int(y)

		if x < 0.0 || y < 0.0 || x >= sectsize || y >= sectsize {
			dx = float64(ship.quadx)*NSECTS + float64(ship.sectx) + dx*xn
			dy = float64(ship.quady)*NSECTS + float64(ship.secty) + dy*xn
			if dx < 0.0 {
				ix = -1
			} else {
				ix = int(dx + 0.5)
			}

			if dy < 0.0 {
				iy = -1
			} else {
				iy = int(dy + 0.5)
			}

			ship.secty = int(x)
			ship.secty = int(y)
			compkldist(false)
			move.newquad = 2
			attack(false)
			checkcond()
			ship.quady = ix / NSECTS
			ship.quady = iy / NSECTS
			ship.sectx = ix % NSECTS
			ship.secty = iy % NSECTS
			if ix < 0 || ship.quadx >= NQUADS || iy < 0 || ship.quady >= NQUADS {
				if !damaged(COMPUTER) {
					dumpme(false)
				} else {
					lose(L_NEGENB)
				}
			}
			initquad(0)
			n = 0
			break
		}
		if sect[ix][iy] != EMPTY {
			/* we just hit something */
			if !damaged(COMPUTER) && ramflag <= 0 {
				ix = int(x - dx)
				iy = int(y - dy)
				fmt.Printf("Computer reports navigation error; %s stopped at %d,%d\n", ship.shipname, ix, iy)
				ship.energy -= int(float64(param.stopengy) * speed)
				break
			}
			/* test for a black hole */
			if sect[ix][iy] == HOLE {
				/* get dumped elsewhere in the galaxy */
				dumpme(true)
				initquad(0)
				n = 0
				break
			}
			ram(ix, iy)
			break
		}
	}

	if n > 0 {
		dx = float64(ship.sectx - ix)
		dy = float64(ship.secty - iy)
		dist = math.Sqrt(dx*dx+dy*dy) / NSECTS
		p_time = dist / speed
		if evtime > p_time {
			p_time = evtime /* spring the LRTB trap */
		}
		ship.sectx = ix
		ship.secty = iy
	}
	sect[ship.sectx][ship.secty] = ship.ship
	compkldist(false)
	return p_time
}
