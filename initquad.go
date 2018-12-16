package main

import "fmt"

func initquad(f int) {
	var (
		rx, ry         int
		nbases, nstars int
		q              *quadrant
		nholes         int
	)

	q = &quad[ship.quadx][ship.quady]

	/* ignored supernova'ed quadrants (this is checked again later anyway */
	if q.stars < 0 {
		return
	}

	etc.enemyCount = q.enemies
	nbases = q.bases
	nstars = q.stars
	nholes = q.holes

	if etc.enemyCount > 0 && !etc.firstContact {
		etc.firstContact = true
		printEnemyGreeting()
	}

	/* have we blundered into a battle zone w/ shields down? */
	if etc.enemyCount > 0 && f == 0 {
		fmt.Printf("Condition RED\n")
		ship.cond = RED
		if !damaged(COMPUTER) {
			shield(1)
		}
	}

	/* clear out the quadrant */
	for i := 0; i < NSECTS; i++ {
		for j := 0; j < NSECTS; j++ {
			sect[i][j] = EMPTY
		}
	}

	/* initialize Enterprise */
	sect[ship.sectx][ship.secty] = ship.ship

	/* initialize enemies */
	for i := 0; i < etc.enemyCount; i++ {
		rx, ry = sector()
		sect[rx][ry] = ENEMY
		etc.enemyList[i].x = rx
		etc.enemyList[i].y = ry
		etc.enemyList[i].power = param.enemyPower
		etc.enemyList[i].srndreq = false
	}

	compkldist(true)

	/* initialize star base */
	if nbases > 0 {
		rx, ry = sector()
		sect[rx][ry] = BASE
		etc.starbase.x = rx
		etc.starbase.y = ry
	}

	/* initialize inhabited starsystem */
	if q.qsystemname != 0 {
		rx, ry = sector()
		sect[rx][ry] = INHABIT
		nstars -= 1
	}

	/* initialize black holes */
	for i := 0; i < nholes; i++ {
		rx, ry = sector()
		sect[rx][ry] = HOLE
	}

	/* initialize stars */
	for i := 0; i < nstars; i++ {
		rx, ry = sector()
		sect[rx][ry] = STAR
	}

	move.newquad = 1
}

func sector() (x, y int) {
	for {
		i := ranf(NSECTS)
		j := ranf(NSECTS)

		if sect[i][j] == EMPTY {
			return i, j
		}
	}
}
