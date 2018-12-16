package main

import (
	"fmt"
	"math"
)

func torped(v int) {
	var (
		ix, iy          int
		x, y, dx, dy    float64
		angle           float64
		course, course2 float64
		bigger          float64
		sectsize        float64
		burst           int
		n               int
	)

	if ship.cloaked {
		fmt.Printf("Federation regulations do not permit attack while cloaked.\n")
		return
	}
	if check_out(TORPED) {
		return
	}
	if ship.torped <= 0 {
		fmt.Printf("All photon torpedos expended\n")
		return
	}

	course = float64(getintpar("Torpedo course"))
	if course < 0 || course > 360 {
		return
	}
	burst = -1

	/* need at least three torpedoes for a burst */
	if ship.torped < 3 {
		fmt.Printf("No-burst mode selected\n")
		burst = 0
	} else if !testnl() {
		/* see if the user wants one */
		c, _ := stdin.ReadByte()
		stdin.UnreadByte()
		if c >= '0' && c <= '9' {
			burst = 1
		}
	}

	if burst < 0 {
		if getynpar("Do you want a burst") {
			burst = 1
		} else {
			burst = 0
		}
	}

	if burst > 0 {
		burst = getintpar("burst angle")
		if burst <= 0 {
			return
		}
		if burst > 15 {
			fmt.Printf("Maximum burst angle is 15 degrees\n")
			return
		}
	}
	sectsize = NSECTS
	n = -1
	if burst > 0 {
		n = 1
		course -= float64(burst)
	}

	for ; n != 0 && n <= 3; n++ {
		/* select a nice random course */
		course2 = course + float64(randcourse(n))
		angle = course2 * 0.0174532925 /* convert to radians */
		dx = -math.Cos(angle)
		dy = math.Sin(angle)
		bigger = math.Abs(dx)
		x = math.Abs(dy)
		if x > bigger {
			bigger = x
		}
		dx /= bigger
		dy /= bigger
		x = float64(ship.sectx) + 0.5
		y = float64(ship.secty) + 0.5
		if ship.cond != DOCKED {
			ship.torped -= 1
		}
		fmt.Printf("Torpedo track")
		if n > 0 {
			fmt.Printf(", torpedo number %d", n)
		}
		fmt.Printf(":\n%6.1f\t%4.1f\n", x, y)
		for {
			x += dx
			y += dy
			ix = int(x)
			iy = int(y)
			if x < 0.0 || x >= sectsize || y < 0.0 || y >= sectsize {
				fmt.Printf("Torpedo missed\n")
				break
			}
			fmt.Printf("%6.1f\t%4.1f\n", x, y)
			switch sect[ix][iy] {
			case EMPTY:
				continue
			case HOLE:
				fmt.Printf("Torpedo disappears into a black hole\n")
			case KLINGON:
				for k := 0; k < etc.nkling; k++ {
					if etc.klingon[k].x != ix || etc.klingon[k].y != iy {
						continue
					}
					etc.klingon[k].power -= 500 + ranf(501)
					if etc.klingon[k].power > 0 {
						fmt.Printf("*** Hit on %s at %d,%d: extensive damages\n", names.enemy, ix, iy)
						break
					}
					killk(ix, iy)
				}
			case STAR:
				nova(ix, iy)
			case INHABIT:
				kills(ix, iy, -1)
			case BASE:
				killb(ship.quadx, ship.quady)
				game.killb += 1
			default:
				fmt.Printf("Unknown object %c at %d,%d destroyed\n", sect[ix][iy], ix, iy)
				sect[ix][iy] = EMPTY
			}
			break
		}
		if damaged(TORPED) || quad[ship.quadx][ship.quady].stars < 0 {
			break
		}
		course += float64(burst)
	}
	move.free = false
}

func randcourse(n int) int {
	var (
		r float64
		d int
	)

	d = int(((franf() + franf()) - 1.0) * 20)
	if math.Abs(float64(d)) > 12 {
		fmt.Printf("Photon tubes misfire")
		if n < 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" on torpedo %d\n", n)
		}
		if ranf(2) != 0 {
			damage(TORPED, 0.2*math.Abs(float64(d))*(franf()+1.0))
		}
		d = int(float64(d) * (1.0 + 2.0*franf()))
	}
	if ship.shldup || ship.cond == DOCKED {
		r = float64(ship.shield)
		r = 1.0 + r/float64(param.shield)
		if ship.cond == DOCKED {
			r = 2.0
		}
		d = int(float64(d) * r)
	}

	return d
}
