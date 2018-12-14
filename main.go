package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	stdin = bufio.NewReader(os.Stdin)
	trace *bool
)

func main() {
	trace = flag.Bool("t", false, "enable tracing output")
	flag.Parse()

	// don't seed randomizer when tracing to ensure the same reproducable steps every time.
	if !*trace {
		rand.Seed(time.Now().UnixNano() ^ int64(os.Getpid()))
	}

	fmt.Printf("\n   * * *   S T A R   T R E K   * * *\n\n")
	fmt.Printf("        __________________           __\n")
	fmt.Printf("        \\_________________|)____.---'--`---.____\n")
	fmt.Printf("                      ||    \\----.________.----/\n")
	fmt.Printf("                      ||     / /    `--'\n")
	fmt.Printf("                    __||____/ /_\n")
	fmt.Printf("                   |___         \\\n")
	fmt.Printf("                       `--------'\n")
	fmt.Printf("\nPress return to continue.\n")

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
