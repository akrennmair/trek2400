package main

var comtab = []cvntab{
	{full: "abandon", funcValue: abandon},
	{abrev: "ca", full: "capture", funcValue: capture},
	{abrev: "cl", full: "cloak", funcValue: shield, intValue: -1},
	{abrev: "c", full: "computer", funcValue: computer},
	{abrev: "da", full: "damages", funcValue: dcrept},
	{full: "destruct", funcValue: destruct},
	{abrev: "do", full: "dock", funcValue: dock},
	{full: "help", funcValue: help},
	{abrev: "i", full: "impulse", funcValue: impulse},
	{abrev: "l", full: "lrscan", funcValue: lrscan},
	{abrev: "m", full: "move", funcValue: dowarp},
	{abrev: "p", full: "phasers", funcValue: phaser},
	{full: "ram", funcValue: dowarp, intValue: 1},
	{full: "dump", funcValue: dumpgame},
	{abrev: "r", full: "rest", funcValue: rest},
	{abrev: "sh", full: "shield", funcValue: shield},
	{abrev: "s", full: "srscan", funcValue: srscan},
	{abrev: "st", full: "status", funcValue: srscan, intValue: -1},
	{full: "terminate", funcValue: myreset},
	{abrev: "t", full: "torpedo", funcValue: torped},
	{abrev: "u", full: "undock", funcValue: undock},
	{abrev: "v", full: "visual", funcValue: visual},
	{abrev: "w", full: "warp", funcValue: setwarp},
}

func play() {
	for {
		move.free = true
		move.time = 0.0
		move.shldchg = false
		move.newquad = 0
		move.resting = false
		skiptonl(0)
		r := getcodpar("\nCommand", comtab)
		r.funcValue(r.intValue)
		events(false)
		attack(false)
		checkcond()
	}
}

func myreset(v int) {
	panic(endofgame{})
}
