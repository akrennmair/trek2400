package main

import (
	"fmt"
	"time"
)

func destruct(v int) {
	var (
		checkpass string
		zap       float64
	)

	if damaged(COMPUTER) {
		out(COMPUTER)
		return
	}

	fmt.Printf("\n\007 --- WORKING ---\007\n")
	time.Sleep(3 * time.Second)
	for i := 10; i > 5; i-- {
		for j := 10; j > i; j-- {
			fmt.Printf("   ")
		}
		fmt.Printf("%d\n", i)
		time.Sleep(1 * time.Second)
	}

	skiptonl()
	checkpass = getstrpar("Enter password verification")
	time.Sleep(2)
	if checkpass != game.passwd {
		fmt.Printf("Self destruct sequence aborted\n")
		return
	}

	fmt.Printf("Password verified; self destruct sequence continues:\n")
	time.Sleep(2 * time.Second)
	for i := 5; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		for j := 5; j > i; j-- {
			fmt.Printf("   ")
		}
		fmt.Printf("%d\n", i)
	}
	time.Sleep(2)
	fmt.Printf("\032\014***** %s destroyed *****\n", ship.shipname)
	game.killed = true
	/* let's see what we can blow up!!!! */
	zap = 20.0 * float64(ship.energy)
	game.deaths += ship.crew
	for i := 0; i < etc.nkling; {
		if float64(etc.klingon[i].power)*etc.klingon[i].dist <= zap {
			killk(etc.klingon[i].x, etc.klingon[i].y)
		} else {
			i++
		}
	}
	/* if we didn't kill the last Klingon (detected by killk), */
	/* then we lose.... */
	lose(L_DSTRCT)
}
