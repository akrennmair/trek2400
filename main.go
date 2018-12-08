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
		func() {
			defer func() {
				// panic/recover is used as a replacement for setjmp/longjmp.
				if e := recover(); e != nil {
					if _, ok := e.(endofgame); !ok {
						panic(e)
					}
				}
			}()
			setup()
			play()
		}()

		if !getynpar("Another game") {
			return
		}
	}
}
