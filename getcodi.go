package main

func getcodi(co *int, di *float64) bool {
	*co = getintpar("Course")

	/* course must be in the interval [0,360] */
	if *co < 0 || *co > 360 {
		return true
	}

	*di = getfltpar("Distance")

	/* distance must be in the interval [0, 15] */
	if *di <= 0.0 || *di > 15.0 {
		return true
	}

	return false
}
