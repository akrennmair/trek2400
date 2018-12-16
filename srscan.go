package main

import (
	"fmt"

	"github.com/fatih/color"
)

func srscan(f int) {
	var (
		statinfo bool
		s        string
		percent  int
		q        *quadrant
		p        *cvntab
	)

	if f >= 0 && check_out(SRSCAN) {
		return
	}
	if f != 0 {
		statinfo = true
	} else {
		if !testnl() {
			etc.statreport = getynpar("status report")
		}
		statinfo = etc.statreport
	}

	if f > 0 {
		etc.statreport = true
	}

	if f >= 0 {
		fmt.Printf("\nShort range sensor scan\n")
		q = &quad[ship.quadx][ship.quady]
		q.scanned = q.enemies*100 + q.bases*10 + q.stars
		fmt.Printf(" ")
		for i := 0; i < NSECTS; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("\n")
	}

	for i := 0; i < NSECTS; i++ {
		if f >= 0 {
			fmt.Printf("%d ", i)
			for j := 0; j < NSECTS; j++ {
				var s string
				c := sect[i][j]
				switch c {
				case BASE:
					s = color.CyanString("%c ", c)
				case STAR:
					s = color.YellowString("%c ", c)
				case ENEMY:
					s = color.RedString("%c ", c)
				case MAINSHIP, QUEENE:
					s = color.HiWhiteString("%c ", c)
				default:
					s = fmt.Sprintf("%c ", sect[i][j])
				}
				fmt.Print(s)
			}
			fmt.Printf("%d", i)
			if statinfo {
				fmt.Printf("   ")
			}
		}
		if statinfo {
			switch i {
			case 0:
				fmt.Printf("stardate      %.2f", now.date)
			case 1:
				fmt.Printf("condition     %s", condColor[ship.cond])
				if ship.cloaked {
					fmt.Printf(", CLOAKED")
				}
			case 2:
				fmt.Printf("position      %d,%d/%d,%d", ship.quadx, ship.quady, ship.sectx, ship.secty)
			case 3:
				fmt.Printf("warp factor   %.1f", ship.warp)
			case 4:
				fmt.Printf("total energy  %d", ship.energy)
			case 5:
				fmt.Printf("torpedoes     %d", ship.torped)
			case 6:
				s = "down"
				if ship.shldup {
					s = "up"
				}
				if damaged(SHIELD) {
					s = "damaged"
				}
				percent = 100 * ship.shield / param.shield
				fmt.Printf("shields       %s, %d%%", s, percent)
			case 7:
				fmt.Printf("%ss left %d", period.enemy, now.enemies)
			case 8:
				fmt.Printf("time left     %.2f", now.time)
			case 9:
				fmt.Printf("life support  ")
				if damaged(LIFESUP) {
					fmt.Printf("damaged, reserves = %.2f", ship.reserves)
				} else {
					fmt.Printf("active")
				}
			}
		}
		fmt.Printf("\n")
	}
	if f < 0 {
		fmt.Printf("current crew  %d\n", ship.crew)
		fmt.Printf("brig space    %d\n", ship.brigfree)
		fmt.Printf("%s power %d\n", period.enemy, param.enemyPower)
		l := game.length - 1
		if game.length > 2 {
			l--
		}

		p = &lentab[l]
		fmt.Printf("Length, Skill %s, ", p.full)
		p = &skilltab[game.skill-1]
		fmt.Printf("%s\n", p.full)
		return
	}
	fmt.Printf("  ")
	for i := 0; i < NSECTS; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	if q.qsystemname&Q_DISTRESSED != 0 {
		fmt.Printf("Distressed ")
	}
	if q.qsystemname != 0 {
		fmt.Printf("Starsystem %s\n", systemname(q))
	}
}

var condColor = []string{"GREEN", "DOCKED", "YELLOW", "RED"}
