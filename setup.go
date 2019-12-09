package main

import (
	"fmt"
	"math"
	"math/rand"
)

var (
	lentab = []cvntab{
		{abrev: "s", full: "short", intValue: 1},
		{abrev: "m", full: "medium", intValue: 2},
		{abrev: "l", full: "long", intValue: 4},
		{abrev: "restart"},
	}

	skilltab = []cvntab{
		{abrev: "n", full: "novice", intValue: 1},
		{abrev: "f", full: "fair", intValue: 2},
		{abrev: "g", full: "good", intValue: 3},
		{abrev: "e", full: "expert", intValue: 4},
		{abrev: "c", full: "commodore", intValue: 5},
		{abrev: "i", full: "impossible", intValue: 6},
	}

	shiptab = []cvntab{
		{abrev: "", full: "ncc1701", intValue: TOS},
		{abrev: "", full: "ent-d", intValue: TNG},
		{abrev: "", full: "ent-e", intValue: ENTE},
	}
)

func setup() {
	var (
		r      *cvntab
		f      float64
		d      int
		klump  int
		ix, iy int
		q      *quadrant
		e      *event
		sum    float64
	)

	period = tosPeriod

	for {
		r = getcodpar("What length game", lentab)
		game.length = r.intValue
		if game.length == 0 {
			if !restartgame() {
				continue
			}
			return
		}
		break
	}

	r = getcodpar("What skill game", skilltab)
	game.skill = r.intValue

	r = getcodpar("Which ship", shiptab)
	game.period = r.intValue
	switch game.period {
	case TOS:
		period = tosPeriod
	case TNG:
		period = tngPeriod
	case ENTE:
		period = entePeriod
		// TODO: implement more time periods.
	}

	game.tourn = false
	game.passwd = getstrpar("Enter a password")
	if game.passwd == "tournament" {
		game.passwd = getstrpar("Enter tournament code")
		game.tourn = true
		d = 0
		for i := 0; i < len(game.passwd); i++ {
			d += int(game.passwd[i]) << uint(i)
		}
		rand.Seed(int64(d))
	}

	printShip()

	now.bases = ranf(6-game.skill) + 2
	param.bases = now.bases

	if game.skill == 6 {
		param.bases, now.bases = 1, 1
	}

	now.time = 6.0*float64(game.length) + 2.0
	param.time = now.time

	i := game.skill
	j := game.length

	now.enemies = int(float64(i*j) * 3.5 * (franf() + 0.75))
	param.enemies = now.enemies

	if param.enemies < i*j*5 {
		now.enemies = i * j * 5
		param.enemies = now.enemies
	}
	if param.enemies <= i { // numerical overflow problem?! research
		now.enemies = 127
		param.enemies = now.enemies
	}

	if game.period == ENTE { // there is only ever one borg cube.
		now.enemies = 1
		param.enemies = now.enemies
	}

	param.energy = period.energy
	ship.energy = param.energy

	param.torped = period.torped
	ship.torped = param.torped
	ship.ship = MAINSHIP
	ship.shipname = period.shipname
	param.shield = period.shield
	ship.shield = param.shield

	param.resource = float64(param.enemies) * param.time
	now.resource = param.resource

	param.crew = period.crew
	ship.crew = param.crew
	param.brigfree = period.brigfree
	ship.brigfree = param.brigfree

	ship.shldup = true
	ship.cond = GREEN
	ship.warp = 5.0
	ship.warp2 = 25.0
	ship.warp3 = 125.0
	ship.sinsbad = false
	ship.cloaked = false
	param.date = float64(ranf(20)+20) * 100
	now.date = param.date

	f = float64(game.skill)
	f = math.Log(f + 0.5)
	param.damfac = map[int]float64{}
	for i := range devices {
		param.damfac[i] = f
	}

	param.damprob = map[int]float64{
		WARP:     70,
		SRSCAN:   110,
		LRSCAN:   110,
		PHASER:   125,
		TORPED:   125,
		IMPULSE:  75,
		SHIELD:   150,
		COMPUTER: 20,
		SSRADIO:  35,
		LIFESUP:  30,
		SINS:     20,
		CLOAK:    50,
		XPORTER:  80,
	}

	sum = 0.0
	for _, v := range param.damprob {
		sum += v
	}
	if sum != 1000 {
		panic(fmt.Sprintf("Device probabilities sum to %f", sum))
	}

	param.dockfac = 0.5
	param.regenfac = (5 - float64(game.skill)) * 0.05
	if param.regenfac < 0.0 {
		param.regenfac = 0.0
	}
	param.warptime = 10
	param.stopengy = 50
	param.shupengy = 40
	i = game.skill
	param.enemyPower = period.initialEnemyPower + period.enemyPowerStep*i
	if i >= 6 {
		param.enemyPower += period.enemyPowerStep
	}
	param.phasfac = 0.8
	param.hitfac = 0.5
	param.enemyCrew = 200
	param.srndrprob = 0.0035
	param.moveprob = map[int]float64{
		KM_OB: 45,
		KM_OA: 40,
		KM_EB: 40,
		KM_EA: 25 + 5*float64(game.skill),
		KM_LB: 0,
		KM_LA: 10 + 10*float64(game.skill),
	}
	param.movefac = map[int]float64{
		KM_OB: 0.09,
		KM_OA: -0.05,
		KM_EB: 0.075,
		KM_EA: -0.06 * float64(game.skill),
		KM_LB: 0.0,
		KM_LA: 0.25,
	}
	param.eventdly = map[int]float64{
		E_SNOVA: 0.5,
		E_LRTB:  25.0,
		E_KATSB: 1.0,
		E_KDESB: 3.0,
		E_ISSUE: 1.0,
		E_SNAP:  0.5,
		E_ENSLV: 0.5,
		E_REPRO: 2.0,
	}
	param.navigcrud = []float64{1.50, 0.75}
	param.cloakenergy = period.cloakenergy
	param.energylow = period.energylow

	for i := 0; i < MAXEVENTS; i++ {
		e = &eventList[i]
		e.date = TOOLARGE
	}

	xsched(E_SNOVA, 1, 0, 0, 0)
	xsched(E_LRTB, param.enemies, 0, 0, 0)
	xsched(E_KATSB, 1, 0, 0, 0)
	xsched(E_ISSUE, 1, 0, 0, 0)
	xsched(E_SNAP, 1, 0, 0, 0)
	ship.sectx = ranf(NSECTS)
	ship.secty = ranf(NSECTS)

	/* setup stars */
	for i := 0; i < NQUADS; i++ {
		for j := 0; j < NQUADS; j++ {
			q = &quad[i][j]
			q.enemies = 0
			q.bases = 0
			q.scanned = -1
			q.stars = ranf(9) + 1
			q.holes = ranf(3) - q.stars/5
		}
	}

	/* select inhabited starsystems */
	for d := 1; d < NINHAB; d++ {
		for {
			i := ranf(NQUADS)
			j := ranf(NQUADS)
			q = &quad[i][j]
			if q.qsystemname == 0 {
				break
			}
		}
		q.qsystemname = d
	}

	/* position starbases */
	for i := 0; i < param.bases; i++ {
		var ix, iy int

		for {
			ix = ranf(NQUADS)
			iy = ranf(NQUADS)
			q = &quad[ix][iy]
			if q.bases > 0 {
				continue
			}
			break
		}

		q.bases = 1
		now.base[i].x = ix
		now.base[i].y = iy
		q.scanned = 1001
		if i == 0 {
			ship.quadx = ix
			ship.quady = iy
		}
	}

	/* position enemies */
	for i := param.enemies; i > 0; {
		klump = ranf(4) + 1
		if klump > i {
			klump = i
		}
		for {
			ix = ranf(NQUADS)
			iy = ranf(NQUADS)
			q = &quad[ix][iy]
			if q.enemies+klump > MAXKLQUAD {
				continue
			}
			q.enemies += klump
			i -= klump
			break
		}
	}

	/* initialize this quadrant */
	fmt.Printf("%d %s%s\n%d starbase", param.enemies, period.enemy, plural(param.enemies), param.bases)
	if param.bases > 1 {
		fmt.Printf("s")
	}
	fmt.Printf(" at %d,%d", now.base[0].x, now.base[0].y)
	for i := 1; i < param.bases; i++ {
		fmt.Printf(", %d,%d", now.base[i].x, now.base[i].y)
	}
	fmt.Printf("\nIt takes %d units to kill a %s\n", param.enemyPower, period.enemy)
	move.free = false
	initquad(0)
	srscan(1)
	attack(false)
}
