package main

import "fmt"

var visdelta = []xy{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
}

func visual(z int) {
	var (
		ix, iy int
		co     int
		c      byte
		v      *xy
	)

	co = getintpar("direction")
	if co < 0 || co > 360 {
		return
	}
	co = (co + 22) / 45

	v = &visdelta[co]

	ix = ship.sectx + v.x
	iy = ship.secty + v.y
	if ix < 0 || ix >= NSECTS || iy < 0 || iy >= NSECTS {
		c = '?'
	} else {
		c = sect[ix][iy]
	}
	fmt.Printf("%d,%d %c ", ix, iy, c)

	co++
	v = &visdelta[co]
	ix = ship.sectx + v.x
	iy = ship.secty + v.y
	if ix < 0 || ix >= NSECTS || iy < 0 || iy >= NSECTS {
		c = '?'
	} else {
		c = sect[ix][iy]
	}
	fmt.Printf("%c ", c)

	co++
	v = &visdelta[co]
	ix = ship.sectx + v.x
	iy = ship.secty + v.y
	if ix < 0 || ix >= NSECTS || iy < 0 || iy >= NSECTS {
		c = '?'
	} else {
		c = sect[ix][iy]
	}

	fmt.Printf("%c %d,%d\n", co, ix, iy)
	move.time = 0.05
	move.free = false
}
