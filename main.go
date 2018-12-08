package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	stdin = bufio.NewReader(os.Stdin)
	trace *bool
)

func main() {
	trace = flag.Bool("t", false, "enable tracing output")
	flag.Parse()

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
