package main

import "fmt"

var udtab = []cvntab{
	{abrev: "u", full: "up", intValue: 1},
	{abrev: "d", full: "down", intValue: 0},
}

func shield(f int) {
	if f > 0 && (ship.shldup || damaged(SRSCAN)) {
		return
	}

	var (
		dev, dev2, dev3 string
		ind             int
		stat            *bool
	)

	if f < 0 {
		/* cloaking device */
		if ship.ship == QUEENE {
			fmt.Printf("Ye Faire Queene does not have the cloaking device.\n")
			return
		}
		dev = "Cloaking device"
		dev2 = "is"
		ind = CLOAK
		dev3 = "it"
		stat = &ship.cloaked
	} else {
		dev = "Shields"
		dev2 = "are"
		dev3 = "them"
		ind = SHIELD
		stat = &ship.shldup
	}

	if damaged(ind) {
		if f <= 0 {
			out(ind)
		}
		return
	}

	if ship.cond == DOCKED {
		fmt.Printf("%s %s down while docked\n", dev, dev2)
		return
	}

	var i bool

	if f <= 0 && !testnl() {
		r := getcodpar("Up or down", udtab)
		i = r.boolValue
	} else {
		var s string
		if *stat {
			s = fmt.Sprintf("%s %s up.  Do you want %s down", dev, dev2, dev3)
		} else {
			s = fmt.Sprintf("%s %s down.  Do you want %s up", dev, dev2, dev3)
		}
		if !getynpar(s) {
			return
		}

		i = !*stat
	}
	if *stat == i {
		fmt.Printf("%s already ", dev)
		if i {
			fmt.Printf("up\n")
		} else {
			fmt.Printf("down\n")
		}
		return
	}

	if i {
		if f >= 0 {
			ship.energy -= param.shupengy
		} else {
			ship.cloakgood = false
		}
	}
	move.free = false
	if f >= 0 {
		move.shldchg = true
	}
	*stat = i
}
