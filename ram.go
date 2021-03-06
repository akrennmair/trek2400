package main

import (
	"fmt"
	"time"
)

func ram(ix, iy int) {
	fmt.Printf("\007RED ALERT\007: collision imminent\n")
	c := sect[ix][iy]
	switch c {
	case KLINGON:
		fmt.Printf("%s rams Klingon at %d,%d\n", ship.shipname, ix, iy)
		killk(ix, iy)

	case STAR, INHABIT:
		fmt.Printf("Yeoman Rand: Captain, isn't it getting hot in here?\n")
		time.Sleep(2)
		fmt.Printf("Spock: Hull temperature approaching 550 Degrees Kelvin.\n") // what's maximum temperature Enterprise can sustain?
		lose(L_STAR)

	case BASE:
		fmt.Printf("You ran into the starbase at %d,%d\n", ix, iy)
		killb(ship.quadx, ship.quady)
		/* don't penalize the captain if it wasn't his fault */
		if !damaged(SINS) {
			game.killb += 1
		}
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("%s heavily damaged\n", ship.shipname)

	/* select the number of deaths to occur */
	i := 10 + ranf(20*game.skill)
	game.deaths += i
	ship.crew -= i
	fmt.Printf("McCoy: Take it easy Jim; we had %d casualties.\n", i)

	/* damage devices with an 80% probability */
	for i := 0; i < len(devices); i++ {
		if ranf(100) < 20 {
			continue
		}
		damage(i, (2.5*(franf()+franf())+1.0)*param.damfac[i])
	}

	/* no chance that your shields remained up in all that */
	ship.shldup = false
}
