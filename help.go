package main

import (
	"fmt"
	"math"
	"time"
)

var cntvect = []string{"first", "second", "third"}

func help(v int) {
	var (
		i       int
		dist, x float64
		dx, dy  int
		j, l    int
	)

	/* check to see if calling for help is reasonable ... */
	if ship.cond == DOCKED {
		fmt.Printf("%s: But Captain, we're already docked\n", names.comms)
		return
	}

	/* or possible */
	if damaged(SSRADIO) {
		out(SSRADIO)
		return
	}
	if now.bases <= 0 {
		fmt.Printf("%s: I'm not getting any response from starbase\n", names.comms)
		return
	}

	/* tut tut, there goes the score */
	game.helps += 1

	/* find the closest base */
	dist = TOOLARGE
	if quad[ship.quadx][ship.quady].bases <= 0 {
		/* there isn't one in this quadrant */
		for i = 0; i < now.bases; i++ {
			/* compute distance */
			dx = now.base[i].x - ship.quadx
			dy = now.base[i].y - ship.quady
			x = float64(dx*dx + dy*dy)
			x = math.Sqrt(x)

			if x < dist {
				dist = x
				l = i
			}
		}

		/* go to that quadrant */
		ship.quadx = now.base[l].x
		ship.quady = now.base[l].y
		initquad(1)
	} else {
		dist = 0.0
	}

	sect[ship.sectx][ship.secty] = EMPTY
	fmt.Printf("Starbase in %d,%d responds\n", ship.quadx, ship.quady)

	/* this next thing acts as a probability that it will work */
	x = math.Pow(1.0-math.Pow(0.94, dist), 0.3333333)

	/* attempt to rematerialize */
	for i = 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		fmt.Printf("%s attept to rematerialize ", cntvect[i])
		if franf() > x {
			/* ok, that's good.  let's see if we can set her down */
			for j = 0; j < 5; j++ {
				dx = etc.starbase.x + ranf(3) - 1
				if dx < 0 || dx >= NSECTS {
					continue
				}
				dy = etc.starbase.y + ranf(3) - 1
				if dy < 0 || dy >= NSECTS || sect[dx][dy] != EMPTY {
					continue
				}
				break
			}
			if j < 5 {
				/* found an empty spot */
				fmt.Printf("succeeds\n")
				ship.sectx = dx
				ship.secty = dy
				sect[dx][dy] = ship.ship
				dock(0)
				compkldist(false)
				return
			}
			/* the starbase must have been surrounded */
		}
		fmt.Printf("fails\n")
	}

	/* one, two, three strikes, you're out */
	lose(L_NOHELP)
}
