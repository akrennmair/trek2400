package main

import "fmt"

func capture(v int) {
	if ship.cloaked {
		fmt.Printf("Ship-ship communications out when cloaked\n")
		return
	}

	if damaged(SSRADIO) {
		out(SSRADIO)
		return
	}

	/* find out if there are any at all */
	if etc.nkling <= 0 {
		fmt.Printf("Uhura: Getting no response, sir\n")
		return
	}

	/* if there is more than one Klingon, find out which one */
	k := selectklingon()
	move.free = false
	move.time = 0.05

	/* check out that Klingon */
	k.srndreq = true
	x := param.klingpwr
	x *= ship.energy
	x /= k.power * etc.nkling
	x = int(float64(x) * param.srndrprob)
	i := x
	if i > ranf(100) {
		/* guess what, he surrendered!!! */
		fmt.Printf("Klingon at %d,%d surrenders\n", k.x, k.y)
		i = ranf(param.klingcrew)
		if i > 0 {
			fmt.Printf("%d klingons commit suicide rather than be taken captive\n", param.klingcrew-i)
		}
		if i > ship.brigfree {
			i = ship.brigfree
		}
		ship.brigfree -= i
		fmt.Printf("%d captives taken", i)
		killk(k.x, k.y)
		return
	}

	/* big surprise, he refuses to surrender */
	fmt.Printf("Fat chance, captain\n")
}

func selectklingon() *kling {
	var i int

	if etc.nkling < 2 {
		i = 0
	} else {
		i = ranf(etc.nkling)
	}

	return &etc.klingon[i]
}
