package main

var comtab = []cvntab{
	{abrev: "s", full: "srscan", funcValue: srscan, intValue: 0},
	{abrev: "st", full: "status", funcValue: srscan, intValue: -1},
}

func play() {
	for {
		move.free = true
		move.time = 0.0
		move.shldchg = false
		move.newquad = false
		move.resting = false
		skiptonl(0)
		r := getcodpar("\nCommand", comtab)
		r.funcValue(r.intValue)
		events(false)
		attack(false)
		checkcond()
	}
}
