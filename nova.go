package main

import "fmt"

func nova(x, y int) {
	if sect[x][y] != STAR || quad[ship.quadx][ship.quady].stars < 0 {
		return
	}
	if ranf(100) < 15 {
		fmt.Printf("Spock: Star at %d,%d failed to nova.\n", x, y)
		return
	}
	if ranf(100) < 5 {
		snova(x, y)
		return
	}
	fmt.Printf("Spock: Star at %d,%d gone nova\n", x, y)
	if ranf(4) != 0 {
		sect[x][y] = EMPTY
	} else {
		sect[x][y] = HOLE
		quad[ship.quadx][ship.quady].holes += 1
	}
	quad[ship.quadx][ship.quady].stars -= 1
	game.kills += 1
	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= NSECTS {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			se := sect[i][j]
			switch se {
			case EMPTY, HOLE:
				// nothing
			case KLINGON:
				killk(i, j)
			case STAR:
				nova(i, j)
			case INHABIT:
				kills(i, j, -1)
			case BASE:
				killb(i, j)
				game.killb += 1
			case ENTERPRISE, QUEENE:
				explosion := 2000
				if ship.shldup {
					if ship.shield >= explosion {
						ship.shield -= explosion
						explosion = 0
					} else {
						explosion -= ship.shield
						ship.shield = 0
					}
				}
				ship.energy -= explosion
				if ship.energy <= 0 {
					lose(L_SUICID)
				}
			default:
				fmt.Printf("Unknown object %c at %d,%d destroyed\n", se, i, j)
				sect[i][j] = EMPTY
			}
		}
	}
}
