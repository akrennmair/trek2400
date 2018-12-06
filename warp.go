package main

func dowarp(fl int) {
	var (
		c int
		d float64
	)

	if getcodi(&c, &d) {
		return
	}
	warp(fl, c, d)
}

func warp(fl, c int, d float64) {
	// TODO: implement
}
