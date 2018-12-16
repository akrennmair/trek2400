package main

import "fmt"

func capture(v int) {
	var (
		i int
		k *enemy
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
	if etc.enemyCount <= 0 {
		fmt.Printf("%s: Getting no response, sir\n", names.comms)
		return
	}

	/* if there is more than one enemy, find out which one */
	k = selectEnemy()
	move.free = false
	move.time = 0.05

	/* check out that enemy */
	k.srndreq = true
	x = float64(param.enemyPower)
	x *= float64(ship.energy)
	x /= float64(k.power * etc.enemyCount)
	x *= param.srndrprob
	i = int(x)
	if i > ranf(100) {
		/* guess what, he surrendered!!! */
		fmt.Printf("%s at %d,%d surrenders\n", names.enemy, k.x, k.y)
		i = ranf(param.enemyCrew)
		if i > 0 {
			fmt.Printf("%d %ss commit suicide rather than be taken captive\n", names.enemy, param.enemyCrew-i)
		}
		if i > ship.brigfree {
			i = ship.brigfree
		}
		ship.brigfree -= i
		fmt.Printf("%d captives taken", i)
		killEnemy(k.x, k.y)
		return
	}

	/* big surprise, he refuses to surrender */
	fmt.Printf("Fat chance, captain\n")
}

func selectEnemy() *enemy {
	var i int

	if etc.enemyCount < 2 {
		i = 0
	} else {
		i = ranf(etc.enemyCount)
	}

	return &etc.enemyList[i]
}
