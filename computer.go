package main

import (
	"fmt"
	"math"
	"os"

	"github.com/fatih/color"
)

var cputab = []cvntab{
	{abrev: "ch", full: "chart", intValue: 1},
	{abrev: "t", full: "trajectory", intValue: 2},
	{abrev: "c", full: "course", intValue: 3},
	{abrev: "m", full: "move", intValue: 3, boolValue: true},
	{abrev: "s", full: "score", intValue: 4},
	{abrev: "p", full: "pheff", intValue: 5},
	{abrev: "w", full: "warpcost", intValue: 6},
	{abrev: "i", full: "impcost", intValue: 7},
	{abrev: "d", full: "idstresslist", intValue: 8},
}

func computer(v int) {
	var (
		ix, iy       int
		j            int
		tqx, tqy     int
		r            *cvntab
		cost         int
		course       int
		dist, p_time float64
		warpfact     float64
		q            *quadrant
		e            *event
	)

	if check_out(COMPUTER) {
		return
	}

	for {
		r = getcodpar("\nRequest", cputab)
		switch r.intValue {

		case 1: /* star chart */
			fmt.Printf("Computer record of galaxy for all long range sensor scans\n\n")
			fmt.Printf("  ")
			/* print top header */
			for i := 0; i < NQUADS; i++ {
				fmt.Printf("-%d- ", i)
			}
			fmt.Printf("\n")
			for i := 0; i < NQUADS; i++ {
				fmt.Printf("%d ", i)
				for j := 0; j < NQUADS; j++ {
					if i == ship.quadx && j == ship.quady {
						fmt.Printf("%s", color.HiWhiteString("o-= "))
						continue
					}
					q = &quad[i][j]
					/* 1000 or 1001 is special case */
					if q.scanned >= 1000 {
						if q.scanned > 1000 {
							fmt.Printf(".%s. ", color.CyanString("1"))
						} else {
							fmt.Printf("%s", color.HiRedString("/// "))
						}
					} else {
						if q.scanned < 0 {
							fmt.Printf("... ")
						} else {
							enemies := q.scanned / 100
							bases := (q.scanned / 10) % 10
							stars := q.scanned % 10

							if enemies > 0 {
								fmt.Printf("%s", color.RedString("%d", enemies))
							} else {
								fmt.Printf(" ")
							}

							if bases > 0 {
								fmt.Printf("%s", color.CyanString("%d", bases))
							} else {
								fmt.Printf(" ")
							}

							fmt.Printf("%s", color.YellowString("%d ", stars))
						}
					}
				}
				fmt.Printf("%d\n", i)
			}
			fmt.Printf("   ")
			/* print bottom footer */
			for i := 0; i < NQUADS; i++ {
				fmt.Printf("-%d- ", i)
			}
			fmt.Printf("\n")

		case 2: /* trajectory */
			if check_out(SRSCAN) {
				break
			}
			if etc.enemyCount <= 0 {
				fmt.Printf("No %ss in this quadrant\n", names.enemy)
			}
			/* for each enemy, give the course & distance */
			for i := 0; i < etc.enemyCount; i++ {
				fmt.Printf("%s at %d,%d", names.enemy, etc.enemyList[i].x, etc.enemyList[i].y)
				course, dist = kalc(ship.quadx, ship.quady, etc.enemyList[i].x, etc.enemyList[i].y)
				prkalc(course, dist)
			}

		case 3: /* course calculation */
			if readdelim('/') {
				tqx = ship.quadx
				tqy = ship.quady
			} else {
				ix = getintpar("Quadrant")
				if ix < 0 || ix >= NSECTS {
					break
				}
				iy = getintpar("q-y")
				if iy < 0 || iy >= NSECTS {
					break
				}
				tqx = ix
				tqy = iy
			}
			ix = getintpar("Sector")
			if ix < 0 || ix >= NSECTS {
				break
			}
			iy = getintpar("s-y")
			if iy < 0 || iy >= NSECTS {
				break
			}
			course, dist = kalc(tqx, tqy, ix, iy)
			if r.boolValue {
				warp(-1, course, dist)
				break
			}
			fmt.Printf("%d,%d/%d,%d to %d,%d/%d,%d", ship.quadx, ship.quady, ship.sectx, ship.secty, tqx, tqy, ix, iy)
			prkalc(course, dist)

		case 4: /* score */
			score()

		case 5: /* phaser effectiveness */
			dist = getfltpar("range")
			if dist < 0.0 {
				break
			}
			dist *= 10.0
			cost = int(math.Pow(0.9, dist)*98.0 + 0.5)
			fmt.Printf("Phasers are %d%% effective at that range\n", cost)

		case 6: /* warp cost (time/energy) */
			dist = getfltpar("distance")
			if dist < 0.0 {
				break
			}
			warpfact = getfltpar("warp factor")
			if warpfact <= 0.0 {
				warpfact = ship.warp
			}
			cost = int((dist + 0.05) * warpfact * warpfact * warpfact)
			p_time = float64(param.warptime) * dist / (warpfact * warpfact)
			fmt.Printf("Warp %.2f distance %.2f cost %.2f stardates %d (%d w/ shlds up) units\n",
				warpfact, dist, p_time, int(cost), int(cost+cost))

		case 7: /* impulse cost */
			dist = getfltpar("distance")
			if dist < 0.0 {
				break
			}
			cost = int(20 + 100*dist)
			p_time = dist / 0.095
			fmt.Printf("Distance %.2f cost %.2f stardates %d units\n", dist, p_time, int(cost))

		case 8: /* distresslist */
			j = 1
			fmt.Printf("\n")
			/* scan the event list */
			for i := 0; i < MAXEVENTS; i++ {
				e = &eventList[i]
				/* ignore hidden entries */
				if e.evcode&E_HIDDEN != 0 {
					continue
				}
				switch e.evcode & E_EVENT {
				case E_KDESB:
					fmt.Printf("%s is attacking starbase in quadrant %d,%d\n", names.enemy, e.x, e.y)
					j = 0
				case E_ENSLV, E_REPRO:
					fmt.Printf("Starsystem %s in quadrant %d,%d is distressed\n", systemnameList[e.systemname], e.x, e.y)
					j = 0
				}
			}
			if j != 0 {
				fmt.Printf("No known distress calls are active\n")
			}
		}

		for {
			i, _ := stdin.ReadByte()
			if i == ';' {
				break
			}
			if i == 0 {
				os.Exit(1)
			}
			if i == '\n' {
				stdin.UnreadByte()
				return
			}
		}
	}
}

func kalc(tqx, tqy, tsx, tsy int) (int, float64) {
	var (
		dx, dy   float64
		quadsize float64
		angle    float64
		course   int
	)

	/* normalize to quadrant distances */
	quadsize = float64(NSECTS)
	dx = (float64(ship.quadx) + float64(ship.sectx)/quadsize) - (float64(tqx) + float64(tsx)/quadsize)
	dy = (float64(tqy) + float64(tsy)/quadsize) - (float64(ship.quady) + float64(ship.secty)/quadsize)

	/* get the angle */
	angle = math.Atan2(dy, dx)
	/* make it 0 -> 2 pi */
	if angle < 0.0 {
		angle += 6.283185307
	}
	/* convert from radians to degrees */
	course = int(angle*57.29577951 + 0.5)
	dx = dx*dx + dy*dy
	return int(course), math.Sqrt(dx)
}

func prkalc(course int, dist float64) {
	fmt.Printf(": course %d  dist %.3f\n", course, dist)
}
