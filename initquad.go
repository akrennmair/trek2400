package main

import "fmt"

func initquad(f int) {
	q := &quad[ship.quadx][ship.quady]

	/* ignored supernova'ed quadrants (this is checked again later anyway */
	if q.stars < 0 {
		return
	}

	etc.nkling = q.klings
	nbases := q.bases
	nstars := q.stars
	nholes := q.holes

	/* have we blundered into a battle zone w/ shields down? */
	if etc.nkling > 0 && f == 0 {
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

	/* initialize Klingons */
	for i := 0; i < etc.nkling; i++ {
		rx, ry := sector()
		sect[rx][ry] = KLINGON
		etc.klingon[i].x = rx
		etc.klingon[i].y = ry
		etc.klingon[i].power = param.klingpwr
		etc.klingon[i].srndreq = false
	}

	compkldist(true)

	/* initialize star base */
	if nbases > 0 {
		rx, ry := sector()
		sect[rx][ry] = BASE
		etc.starbase.x = rx
		etc.starbase.y = ry
	}

	/* initialize inhabited starsystem */
	if q.qsystemname != 0 {
		rx, ry := sector()
		sect[rx][ry] = INHABIT
		nstars -= 1
	}

	/* initialize black holes */
	for i := 0; i < nholes; i++ {
		rx, ry := sector()
		sect[rx][ry] = HOLE
	}

	/* initialize stars */
	for i := 0; i < nstars; i++ {
		rx, ry := sector()
		sect[rx][ry] = STAR
	}

	move.newquad = true
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
