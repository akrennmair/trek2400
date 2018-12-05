package main

import (
	"fmt"
	"math"
)

func events(timeWarp bool) {

	/* if nothing happened, just allow for any Klingons killed */
	if move.time <= 0.0 {
		now.time = now.resource / float64(now.klings)
		return
	}

	/* indicate that the cloaking device is now working */
	ship.cloakgood = true

	/* idate is the initial date */
	idate := now.date

	/* schedule attacks if resting too long */
	if move.time > 0.5 && move.resting {
		schedule(E_ATTACK, 0.5, 0, 0, 0)
	}

	/* scan the event list */
	for {
		restcancel := 0
		evnum := -1
		/* xdate is the date of the current event */
		xdate := idate + move.time

		var e, ev *event

		/* find the first event that has happened */
		for i := 0; i < MAXEVENTS; i++ {
			e = &eventList[i]
			if e.evcode == 0 || (e.evcode&E_GHOST) != 0 {
				continue
			}
			if e.date < xdate {
				xdate = e.date
				ev = e
				evnum = i
			}
		}
		e = ev

		/* find the time between events */
		rtime := xdate - now.date

		/* decrement the magic "Federation Resources" pseudo-variable */
		now.resource -= float64(now.klings) * rtime
		/* and recompute the time left */
		now.time = now.resource / float64(now.klings)

		/* move us up to the next date */
		now.date = xdate

		/* check for out of time */
		if now.time <= 0.0 {
			lose(L_NOTIME)
		}

		/* if evnum < 0, no events occurred  */
		if evnum < 0 {
			break
		}

		/* otherwise one did.  Find out what it is */
		switch e.evcode & E_EVENT {
		case E_SNOVA: /* supernova */
			/* cause the supernova to happen */
			snova(-1, 0)
			/* and schedule the next one */
			xresched(e, E_SNOVA, 1)

		case E_LRTB: /* long range tractor beam */
			/* schedule the next one */
			xresched(e, E_LRTB, now.klings)
			/* LRTB cannot occur if we are docked */
			if ship.cond != DOCKED {
				/* pick a new quadrant */
				i := ranf(now.klings) + 1
				var ix, iy int
				for ix = 0; i < NQUADS; ix++ {
					for iy = 0; iy < NQUADS; iy++ {
						q := &quad[ix][iy]
						if q.stars >= 0 {
							i -= q.klings
							if i <= 0 {
								break
							}
						}
					}
					if i <= 0 {
						break
					}
				}

				/* test for LRTB to same quadrant */
				if ship.quadx == ix && ship.quady == iy {
					break
				}

				/* nope, dump him in the new quadrant */
				ship.quadx = ix
				ship.quady = iy
				fmt.Printf("\n%s caught in long range tractor beam\n", ship.shipname)
				fmt.Printf("*** Pulled to quadrant %d,%d\n", ship.quadx, ship.quady)
				ship.sectx = ranf(NSECTS)
				ship.secty = ranf(NSECTS)
				initquad(0)
				/* truncate the move time */
				move.time = xdate - idate
			}

		case E_KATSB: /* Klingon attacks starbase */
			/* if out of bases, forget it */
			if now.bases <= 0 {
				unschedule(e)
				break
			}

			var i, ix, iy int
			for ; i < now.bases; i++ {
				ix := now.base[i].x
				iy := now.base[i].y
				/* see if a Klingon exists in this quadrant */
				q := &quad[ix][iy]
				if q.klings <= 0 {
					continue
				}

				/* see if already distressed */
				j := 0
				for ; j < MAXEVENTS; j++ {
					e := &eventList[j]
					if (e.evcode & E_EVENT) != E_KDESB {
						continue
					}
					if e.x == ix && e.y == iy {
						break
					}
				}

				if j < MAXEVENTS {
					continue
				}

				break
			}
			e = ev
			if i >= now.bases {
				/* not now; wait a while and see if some Klingons move in */
				reschedule(e, 0.5+3.0*franf())
				break
			}
			/* schedule a new attack, and a destruction of the base */
			xresched(e, E_KATSB, 1)
			e = xsched(E_KDESB, 1, ix, iy, 0)

			/* report it if we can */
			if !damaged(SSRADIO) {
				fmt.Printf("\nUhura:  Captain, we have received a distress signal\n")
				fmt.Printf("  from the starbase in quadrant %d,%d.\n", ix, iy)
				restcancel++
			} else {
				/* SSRADIO out, make it so we can't see the distress call */
				/* but it's still there!!! */
				e.evcode |= E_HIDDEN
			}

		case E_KDESB: /* Klingon destroys starbase */
			unschedule(e)
			q := &quad[e.x][e.y]
			/* if the base has mysteriously gone away, or if the Klingon
			   got tired and went home, ignore this event */
			if q.bases <= 0 || q.klings <= 0 {
				break
			}
			if e.x == ship.quadx && e.y == ship.quady {
				/* yep, kill one in this quadrant */
				fmt.Printf("\nSpock: ")
				killb(ship.quadx, ship.quady)
			} else {
				/* kill one in some other quadrant */
				killb(e.x, e.y)
			}

		case E_ISSUE: /* issue a distress call */
			xresched(e, E_ISSUE, 1)
			/* if we already have too many, throw this one away */
			if ship.distressed >= MAXDISTR {
				break
			}
			var (
				i, ix, iy int
				q         *quadrant
			)
			for ; i < 100; i++ {
				ix = ranf(NQUADS)
				iy = ranf(NQUADS)
				q = &quad[ix][iy]
				/* need a quadrant which is not the current one,
				   which has some stars which are inhabited and
				   not already under attack, which is not
				   supernova'ed, and which has some Klingons in it */
				if !((ix == ship.quadx && iy == ship.quady) || q.stars < 0 || (q.qsystemname&Q_DISTRESSED) != 0 || (q.qsystemname&Q_SYSTEM) == 0 || q.klings < 0) {
					break
				}
			}

			if i >= 100 {
				/* can't seem to find one; ignore this call */
				break
			}

			/* got one!!  Schedule its enslavement */
			ship.distressed++
			e = xsched(E_ENSLV, 1, ix, iy, q.qsystemname)
			//q.qsystemname = (e - eventList) | Q_DISTRESSED // WTF?!
			q.qsystemname = e.systemname | Q_DISTRESSED // I don't know whether this is correct. The line above is the original code.

			/* tell the captain about it if we can */
			if !damaged(SSRADIO) {
				fmt.Printf("\nUhura: Captain, starsystem %s in quadrant %d,%d is under attack\n",
					systemnameList[e.systemname], ix, iy)
				restcancel++
			} else {
				/* if we can't tell him, make it invisible */
				e.evcode |= E_HIDDEN
			}

		case E_ENSLV: /* starsystem is enslaved */
			unschedule(e)
			/* see if current distress call still active */
			q := &quad[e.x][e.y]
			if q.klings <= 0 {
				/* no Klingons, clean up */
				/* restore the system name */
				q.qsystemname = e.systemname
				break
			}

			/* play stork and schedule the first baby */
			e = schedule(E_REPRO, param.eventdly[E_REPRO]*franf(), e.x, e.y, e.systemname)

			/* report the disaster if we can */
			if !damaged(SSRADIO) {
				fmt.Printf("\nUhura:  We've lost contact with starsystem %s\n",
					systemnameList[e.systemname])
				fmt.Printf("  in quadrant %d,%d.\n", e.x, e.y)
			} else {
				e.evcode |= E_HIDDEN
			}

		case E_REPRO: /* Klingon reproduces */
			q := &quad[e.x][e.y]
			if q.klings >= 0 {
				unschedule(e)
				q.qsystemname = e.systemname
				break
			}
			xresched(e, E_REPRO, 1)
			/* reproduce one Klingon */
			ix := e.x
			iy := e.y
			if now.klings == 127 {
				break /* full right now */
			}
			if q.klings >= MAXKLQUAD {
				var i, j int
				/* this quadrant not ok, pick an adjacent one */
				for i = ix - 1; i <= ix+1; i++ {
					if i < 0 || i >= NQUADS {
						continue
					}
					j = iy - 1
					for ; j <= iy+1; j++ {
						if j < 0 || j >= NQUADS {
							continue
						}
						q := &quad[i][j]
						/* check for this quad ok (not full & no snova) */
						if q.klings >= MAXKLQUAD || q.stars < 0 {
							continue
						}
						break
					}
					if j <= iy+1 {
						break
					}
				}
				if j > iy+1 {
					/* cannot create another yet */
					break
				}
				ix = i
				iy = j
			}
			/* deliver the child */
			q.klings++
			now.klings++
			if ix == ship.quadx && iy == ship.quady {
				ix, iy = sector()
				sect[ix][iy] = KLINGON
				k := &etc.klingon[etc.nkling]
				etc.nkling++
				k.x = ix
				k.y = iy
				k.power = param.klingpwr
				k.srndreq = false
				compkldist(etc.klingon[0].dist == etc.klingon[0].avgdist)
			}

			/* recompute time left */
			now.time = now.resource / float64(now.klings)

		case E_SNAP: /* take a snapshot of the galaxy */
			// TODO: not implemented yet.

		case E_ATTACK: /* Klingons attack during rest period */
			if !move.resting {
				unschedule(e)
				break
			}
			attack(true)
			reschedule(e, 0.5)

		case E_FIXDV:
			i := e.systemname
			unschedule(e)

			/* de-damage the device */
			fmt.Printf("%s reports repair work on the %s finished.\n", devices[i].person, devices[i].name)

			/* handle special processing upon fix */
			switch i {
			case LIFESUP:
				ship.reserves = param.reserves
			case SINS:
				if ship.cond == DOCKED {
					break
				}
				fmt.Printf("Spock has tried to recalibrate your Space Internal Navigation System,\n")
				fmt.Printf("  but he has no standard base to calibrate to.  Suggest you get\n")
				fmt.Printf("  to a starbase immediately so that you can properly recalibrate.\n")
				ship.sinsbad = true
			case SSRADIO:
				restcancel = dumpssradio()
			}
		} /* switch */

		if restcancel > 0 && move.resting && getynpar("Spock: Shall we cancel our rest period") {
			move.time = xdate - idate
		}
	}

	if e := now.eventptr[E_ATTACK]; e != nil {
		unschedule(e)
	}

	if !timeWarp {
		/* eat up energy if cloaked */
		if ship.cloaked {
			ship.energy -= int(float64(param.cloakenergy) * move.time)
		}

		/* regenerate resources */
		rtime := 1.0 - math.Exp(-param.regenfac*move.time)
		ship.shield += int(float64(param.shield-ship.shield) * rtime)
		ship.energy += int(float64(param.energy-ship.energy) * rtime)

		/* decrement life support reserves */
		if damaged(LIFESUP) && ship.cond != DOCKED {
			ship.reserves -= move.time
		}
	}
}
