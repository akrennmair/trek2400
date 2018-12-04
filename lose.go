package main

import (
	"fmt"
	"time"
)

var losemsg = map[int]string{
	L_NOTIME:   "You ran out of time",
	L_NOENGY:   "You ran out of energy",
	L_DSTRYD:   "You have been destroyed",
	L_NEGENB:   "You ran into the negative energy barrier",
	L_SUICID:   "You destroyed yourself by nova'ing that star",
	L_SNOVA:    "You have been caught in a supernova",
	L_NOLIFE:   "You just suffocated in outer space",
	L_NOHELP:   "You could not be rematerialized",
	L_TOOFAST:  "\n\032\014 ***\007 Ship's hull has imploded\007 ***",
	L_STAR:     "You have burned up in a star",
	L_DSTRCT:   "Well, you destroyed yourself, but it didn't do any good",
	L_CAPTURED: "You have been captured by Klingons and mercilessly tortured",
	L_NOCREW:   "Your last crew member died",
}

func lose(why int) {
	game.killed = true
	time.Sleep(1 * time.Second)
	fmt.Printf("\n%s\n", losemsg[why])
	if why == L_NOTIME {
		game.killed = false
	}

	move.endgame = -1
	score()
	skiptonl(0)
	// TODO: how to implement longjmp?
}
