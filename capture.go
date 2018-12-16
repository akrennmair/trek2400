package main

import "fmt"

func capture(v int) {
	var (
		i int
		k *kling
		x float64
	)

	/* check for not cloaked */
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
		fmt.Printf("%s: Getting no response, sir\n", names.comms)
		return
	}

	/* if there is more than one Klingon, find out which one */
	k = selectklingon()
	move.free = false
	move.time = 0.05

	/* check out that Klingon */
	k.srndreq = true
	x = float64(param.klingpwr)
	x *= float64(ship.energy)
	x /= float64(k.power * etc.nkling)
	x *= param.srndrprob
	i = int(x)
	if i > ranf(100) {
		/* guess what, he surrendered!!! */
		fmt.Printf("%s at %d,%d surrenders\n", names.enemy, k.x, k.y)
		i = ranf(param.klingcrew)
		if i > 0 {
			fmt.Printf("%d %ss commit suicide rather than be taken captive\n", names.enemy, param.klingcrew-i)
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
