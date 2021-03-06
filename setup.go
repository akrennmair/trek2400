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

	now.bases = ranf(6-game.skill) + 2
	param.bases = now.bases

	if game.skill == 6 {
		param.bases, now.bases = 1, 1
	}

	now.time = 6.0*float64(game.length) + 2.0
	param.time = now.time

	i := game.skill
	j := game.length

	now.klings = int(float64(i*j) * 3.5 * (franf() + 0.75))
	param.klings = now.klings

	if param.klings < i*j*5 {
		now.klings = i * j * 5
		param.klings = now.klings
	}
	if param.klings <= i { // numerical overflow problem?! research
		now.klings = 127
		param.klings = now.klings
	}

	param.energy, ship.energy = 5000, 5000
	param.torped, ship.torped = 10, 10
	ship.ship = ENTERPRISE
	ship.shipname = "Enterprise"
	param.shield, ship.shield = 1500, 1500

	param.resource = float64(param.klings) * param.time
	now.resource = param.resource

	param.crew, ship.crew = 387, 387
	param.brigfree, ship.brigfree = 400, 400

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
	param.klingpwr = 100 + 150*i
	if i >= 6 {
		param.klingpwr += 150
	}
	param.phasfac = 0.8
	param.hitfac = 0.5
	param.klingcrew = 200
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
	param.cloakenergy = 1000
	param.energylow = 1000

	for i := 0; i < MAXEVENTS; i++ {
		e = &eventList[i]
		e.date = TOOLARGE
	}

	xsched(E_SNOVA, 1, 0, 0, 0)
	xsched(E_LRTB, param.klings, 0, 0, 0)
	xsched(E_KATSB, 1, 0, 0, 0)
	xsched(E_ISSUE, 1, 0, 0, 0)
	xsched(E_SNAP, 1, 0, 0, 0)
	ship.sectx = ranf(NSECTS)
	ship.secty = ranf(NSECTS)

	/* setup stars */
	for i := 0; i < NQUADS; i++ {
		for j := 0; j < NQUADS; j++ {
			q = &quad[i][j]
			q.klings = 0
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

	/* position klingons */
	for i := param.klings; i > 0; {
		klump = ranf(4) + 1
		if klump > i {
			klump = i
		}
		for {
			ix = ranf(NQUADS)
			iy = ranf(NQUADS)
			q = &quad[ix][iy]
			if q.klings+klump > MAXKLQUAD {
				continue
			}
			q.klings += klump
			i -= klump
			break
		}
	}

	/* initialize this quadrant */
	fmt.Printf("%d Klingons\n%d starbase", param.klings, param.bases)
	if param.bases > 1 {
		fmt.Printf("s")
	}
	fmt.Printf(" at %d,%d", now.base[0].x, now.base[0].y)
	for i := 1; i < param.bases; i++ {
		fmt.Printf(", %d,%d", now.base[i].x, now.base[i].y)
	}
	fmt.Printf("\nIt takes %d units to kill a Klingon\n", param.klingpwr)
	move.free = false
	initquad(0)
	srscan(1)
	attack(false)
}
