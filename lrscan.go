package main

import (
	"fmt"

	"github.com/fatih/color"
)

func lrscan(v int) {
	var q *quadrant

	if check_out(LRSCAN) {
		return
	}

	fmt.Printf("Long range scan for quadrant %d,%d\n\n", ship.quadx, ship.quady)

	/* print the header on top */
	for j := ship.quady - 1; j <= ship.quady+1; j++ {
		if j < 0 || j >= NQUADS {
			fmt.Printf("      ")
		} else {
			fmt.Printf("     %1d", j)
		}
	}

	/* scan the quadrants */
	for i := ship.quadx - 1; i <= ship.quadx+1; i++ {
		fmt.Printf("\n  -------------------\n")
		if i < 0 || i >= NQUADS {
			fmt.Printf("  !  *  !  *  !  *  !")
			continue
		}

		/* print the left hand margin */
		fmt.Printf("%1d !", i)
		for j := ship.quady - 1; j <= ship.quady+1; j++ {
			if j < 0 || j >= NQUADS {
				/* negative energy barrier again */
				fmt.Printf("  *  !")
				continue
			}
			q = &quad[i][j]
			if q.stars < 0 {
				/* supernova */
				fmt.Printf(" %s !", color.HiRedString("///"))
				q.scanned = 1000
				continue
			}
			q.scanned = q.klings*100 + q.bases*10 + q.stars
			fmt.Printf(" ")
			if q.klings > 0 {
				fmt.Printf("%s", color.RedString("%d", q.klings))
			} else {
				fmt.Printf(" ")
			}

			if q.bases > 0 {
				fmt.Printf("%s", color.CyanString("%d", q.bases))
			} else {
				fmt.Printf(" ")
			}

			fmt.Printf("%s !", color.YellowString("%d", q.stars))
		}
	}
	fmt.Printf("\n  -------------------\n")
}
