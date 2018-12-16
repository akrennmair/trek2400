package main

func checkcond() {
	if ship.reserves < 0.0 {
		lose(L_NOLIFE)
	}

	if ship.energy <= 0 {
		lose(L_NOENGY)
	}

	if ship.crew <= 0 {
		lose(L_NOCREW)
	}

	if etc.enemyCount < 0 {
		return
	}

	if quad[ship.quadx][ship.quady].stars < 0 {
		autover()
	}

	if quad[ship.quadx][ship.quady].stars < 0 {
		lose(L_SNOVA)
	}

	if etc.enemyCount <= 0 {
		killd(ship.quadx, ship.quady, 1)
	}

	if ship.cond == DOCKED {
		return
	}

	if etc.enemyCount > 0 {
		ship.cond = RED
		return
	}

	if ship.energy < param.energylow {
		ship.cond = YELLOW
		return
	}

	ship.cond = GREEN
}
