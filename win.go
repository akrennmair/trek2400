package main

import (
	"fmt"
	"time"
)

func win() {
	time.Sleep(1 * time.Second)
	fmt.Printf("\nCongratulations, you have saved the Federation\n")
	move.endgame = 1

	/* print and return the score */
	s := score()

	/* decide if she gets a promotion */
	if game.helps == 0 && game.killb == 0 && game.killinhab == 0 && 5*game.kills+game.deaths < 100 && s >= 1000 && ship.ship == ENTERPRISE {
		fmt.Printf("In fact, you are promoted one step in rank,\n")
		if game.skill >= 6 {
			fmt.Printf("to the exalted rank of Commodore Emeritus\n")
		} else {
			fmt.Printf("from %s to %s\n", skilltab[game.skill-1].full, skilltab[game.skill].full)
		}
	}

	skiptonl(0)
	panic(endofgame{})
}
