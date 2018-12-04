package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	stdin = bufio.NewReader(os.Stdin)
)

func main() {
	fmt.Printf("\n   * * *   S T A R   T R E K   * * *\n\nPress return to continue.\n")

	for {
		setup()
		play()

		if !getynpar("Another game") {
			return
		}
	}
}
