package main

import (
	"fmt"
	"math"
)

const (
	ALPHA   = 3.0    /* spread */
	BETA    = 3.0    /* franf() */
	GAMMA   = 0.30   /* cos(angle) */
	EPSILON = 150.0  /* dist ** 2 */
	OMEGA   = 10.596 /* overall scaling factor */
)

var matab = []cvntab{
	{abrev: "m", full: "manual", boolValue: true},
	{abrev: "a", full: "automatic", boolValue: false},
}

type banks struct {
	units  int
	angle  float64
	spread float64
}

func phaser(v int) {
	var (
		i, j                    int
		k                       *kling
		dx, dy                  float64
		anglefactor, distfactor float64
		b                       *banks
		flag, extra             int
		manual                  bool
		hit                     int
		tot                     float64
		n                       int
		hitreqd                 [NBANKS]int
		bank                    [NBANKS]banks
		ptr                     *cvntab
	)

	if ship.cond == DOCKED {
		fmt.Printf("Phasers cannot fire through starbase shields\n")
		return
	}

	if damaged(PHASER) {
		out(PHASER)
		return
	}

	if ship.shldup {
		fmt.Printf("Sulu: Captain, we cannot fire through shields.\n")
		return
	}

	if ship.cloaked {
		fmt.Printf("Sulu: Captain, surely you must realize that we cannot fire\n")
		fmt.Printf("  phasers with the cloaking device up.\n")
		return
	}

	/* decide if we want manual or automatic mode */
	manual = false
	if testnl() {
		if damaged(COMPUTER) {
			fmt.Printf("%s", devices[COMPUTER].name)
			manual = true
		} else if damaged(SRSCAN) {
			fmt.Printf("%s", devices[SRSCAN].name)
			manual = true
		}
		if manual {
			fmt.Printf(" damaged, manual mode selected\n")
		}
	}

	if !manual {
		ptr = getcodpar("Manual or automatic", matab)
		manual = ptr.boolValue
	}

	if !manual && damaged(COMPUTER) {
		fmt.Printf("Computer damaged, manual selected\n")
		skiptonl()
		manual = true
	}

	flag = 1
	if manual {
		for flag != 0 {
			fmt.Printf("%d units available\n", ship.energy)
			extra = 0
			flag = 0
			for i := 0; i < NBANKS; i++ {
				b = &bank[i]
				fmt.Printf("\nBank %d:\n", i)
				hit = getintpar("units")
				if hit < 0 {
					return
				}
				extra += hit
				if extra > ship.energy {
					fmt.Printf("available energy exceeded.  ")
					skiptonl()
					flag++
					break
				}
				b.units = hit
				course := getintpar("course")
				if course < 0 || course > 360 {
					return
				}
				b.angle = float64(course) * 0.0174532925
				b.spread = getfltpar("spread")
				if b.spread < 0 || b.spread > 1 {
					return
				}
			}
			ship.energy -= extra
		}
		extra = 0
	} else {
		/* automatic distribution of power */
		if etc.nkling <= 0 {
			fmt.Printf("Sulu: But there are no Klingons in this quadrant\n")
			return
		}
		fmt.Printf("Phasers locked on target.  ")
		for flag != 0 {
			fmt.Printf("%d units available\n", ship.energy)
			hit = getintpar("Units to fire")
			if hit <= 0 {
				return
			}
			if hit > ship.energy {
				fmt.Printf("available energy exceeded.  ")
				skiptonl()
				continue
			}
			flag = 0
			ship.energy -= hit
			extra = hit
			n = etc.nkling
			if n > NBANKS {
				n = NBANKS
			}
			tot = float64(n * (n + 1) / 2)
			for i = 0; i < n; i++ {
				k = &etc.klingon[i]
				b = &bank[i]
				distfactor = k.dist
				anglefactor = ALPHA * BETA * OMEGA / (distfactor*distfactor + EPSILON)
				anglefactor *= GAMMA
				distfactor = float64(k.power)
				distfactor /= anglefactor
				hitreqd[i] = int(distfactor + 0.5)
				dx = float64(ship.sectx - k.x)
				dy = float64(k.y - ship.secty)
				b.angle = math.Atan2(dy, dx)
				b.spread = 0.0
				b.units = ((n - i) / int(tot)) * extra
				extra -= b.units
				hit = b.units - hitreqd[i]
				if hit > 0 {
					extra += hit
					b.units -= hit
				}
			}

			/* give out any extra energy we might have around */
			if extra > 0 {
				for i := 0; i < n; i++ {
					b = &bank[i]
					hit = hitreqd[i] - b.units
					if hit <= 0 {
						continue
					}
					if hit >= extra {
						b.units += extra
						extra = 0
						break
					}
					b.units = hitreqd[i]
					extra -= hit
				}
				if extra > 0 {
					fmt.Printf("%d units overkill\n", extra)
				}
			}
		}
	}

	kidx := 0
	move.free = false
	for i = 0; i < NBANKS; i++ {
		b = &bank[i]
		if b.units <= 0 {
			continue
		}
		fmt.Printf("\nPhaser bank %d fires:\n", i)
		n = etc.nkling
		k = &etc.klingon[kidx]
		for j = 0; j < n; j++ {
			if b.units <= 0 {
				break
			}

			// TODO: copy documentation
			distfactor = BETA + franf()
			distfactor *= ALPHA + b.spread
			distfactor *= OMEGA
			anglefactor = k.dist
			distfactor /= anglefactor*anglefactor + EPSILON
			distfactor *= float64(b.units)
			dx = float64(ship.sectx - k.x)
			dy = float64(k.y - ship.secty)
			anglefactor = math.Atan2(dy, dx) - b.angle
			anglefactor = math.Cos((anglefactor * b.spread) + GAMMA)
			if anglefactor < 0.0 {
				kidx++
				continue
			}
			hit = int(anglefactor*distfactor + 0.5)
			k.power -= hit
			fmt.Printf("%d unit hit on Klingon", hit)
			if !damaged(SRSCAN) {
				fmt.Printf(" at %d,%d", k.x, k.y)
			}
			fmt.Printf("\n")
			b.units -= hit
			if k.power <= 0 {
				killk(k.x, k.y)
				continue
			}
			kidx++
		}
	}

	for i = 0; i < NBANKS; i++ {
		extra += bank[i].units
	}
	if extra > 0 {
		fmt.Printf("\n%d units expended on empty space\n", extra)
	}
}
