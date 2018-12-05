package main

import "fmt"

func killb(qx, qy int) {
	// TODO: implement
}

func killd(x, y, f int) {
	q := &quad[x][y]

	for i := 0; i < MAXEVENTS; i++ {
		e := &eventList[i]
		if e.x != x || e.y != y {
			continue
		}
		switch e.evcode {
		case E_KDESB:
			if f != 0 {
				fmt.Printf("Distress call for starbase in %d,%d nullified\n", x, y)
				unschedule(e)
			}
			break
		case E_ENSLV, E_REPRO:
			if f != 0 {
				fmt.Printf("Distress call for %s in quadrant %d,%d nullified\n",
					systemname[e.systemname], x, y)
				q.qsystemname = e.systemname
				unschedule(e)
			} else {
				e.evcode |= E_GHOST
			}
		}
	}
}

var systemname = []string{
	"ERROR",
	"Talos IV",
	"Rigel III",
	"Deneb VII",
	"Canopus V",
	"Icarus I",
	"Prometheus II",
	"Omega VII",
	"Elysium I",
	"Scalos IV",
	"Procyon IV",
	"Arachnid I",
	"Argo VIII",
	"Triad III",
	"Echo IV",
	"Nimrod III",
	"Nemisis IV",
	"Centarurus I",
	"Kronos III",
	"Spectros V",
	"Beta III",
	"Gamma Tranguli VI",
	"Pyris III",
	"Triachus",
	"Marcus XII",
	"Kaland",
	"Ardana",
	"Stratos",
	"Eden",
	"Arrikis",
	"Epsilon Eridani IV",
	"Exo III",
}
