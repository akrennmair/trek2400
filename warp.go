package main

import (
	"fmt"
	"time"
)

func dowarp(fl int) {
	var (
		c int
		d float64
	)

	if getcodi(&c, &d) {
		return
	}
	warp(fl, c, d)
}

func warp(fl, c int, d float64) {
	var (
		course  int
		power   float64
		dist    float64
		p_time  float64
		speed   float64
		frac    float64
		percent int
	)

	if ship.cond == DOCKED {
		fmt.Printf("%s is docked\n", ship.shipname)
		return
	}

	if damaged(WARP) {
		out(WARP)
		return
	}

	course = c
	dist = d

	/* check to see that we are not using an absurd amount of power */
	power = (dist + 0.05) * ship.warp3
	percent = int(100*power/float64(ship.energy) + 0.5)
	if percent >= 85 {
		fmt.Printf("%s: That would consume %d%% of our remaining energy.\n", names.engineer, percent)
		if !getynpar("Are you sure that is wise") {
			return
		}
	}

	/* compute the speed we will move at, and the time it will take */
	speed = ship.warp2 / float64(param.warptime)
	p_time = dist / speed

	/* check to see that that value is not ridiculous */
	percent = int(100*p_time/now.time + 0.5)
	if percent >= 85 {
		fmt.Printf("%s: That would take %d%% of our remaining time.\n", names.firstOfficer, percent)
		if !getynpar("Are you sure that is wise") {
			return
		}
	}

	/* compute how far we will go if we get damages */
	if ship.warp > 6.0 && ranf(100) < int(20+15*(ship.warp-6.0)) {
		frac = franf()
		dist *= frac
		p_time *= frac
		damage(WARP, (frac+1.0)*ship.warp*(franf()+0.25)*0.20)
	}

	/* do the move */
	move.time = domove(fl, course, p_time, speed)

	/* see how far we actually went, and decrement energy appropriately */
	dist = move.time * speed
	b := 0.0
	if ship.shldup {
		b = 1.0
	}
	ship.energy -= int(dist * ship.warp3 * (b + 1))

	/* test for bizarre events */
	if ship.warp <= 9.0 {
		return
	}
	fmt.Printf("\n\n  ___ Speed exceeding warp nine ___\n\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("Ship's safety systems malfunction\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("Crew experiencing extreme sensory distortion\n")
	time.Sleep(4 * time.Second)
	if float64(ranf(100)) >= 100*dist {
		fmt.Printf("Equilibrium restored -- all systems normal\n")
		return
	}

	/* select a bizzare thing to happen to us */
	percent = ranf(100)
	if percent < 70 {
		/* time warp */
		if percent < 35 || !game.snap {
			/* positive time warp */
			p_time = (ship.warp - 8.0) * dist * (franf() + 1.0)
			now.date += p_time
			fmt.Printf("Positive time portal entered -- it is now Stardate %.2f\n", now.date)
			for i := 0; i < MAXEVENTS; i++ {
				percent = eventList[i].evcode
				if percent == E_FIXDV || percent == E_LRTB {
					eventList[i].date += p_time
				}
			}
			return
		}

		/* s/he got lucky: a negative time portal */
		p_time = now.date
		quad = etc.snapshot.quad
		eventList = etc.snapshot.event
		now = etc.snapshot.now
		fmt.Printf("Negative time portal entered -- it is now Stardate %.2f\n", now.date)
		for i := 0; i < MAXEVENTS; i++ {
			if eventList[i].evcode == E_FIXDV {
				reschedule(&eventList[i], eventList[i].date-p_time)
			}
		}
		return
	}

	/* test for just a lot of damage */
	if percent < 80 {
		lose(L_TOOFAST)
	}
	fmt.Printf("Equilibrium restored -- extreme damage occurred to ship systems\n")
	for i := 0; i < len(devices); i++ {
		damage(i, (3.0*(franf()+franf())+1.0)*param.damfac[i])
	}
	ship.shldup = false
}
