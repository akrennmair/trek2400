package main

import "fmt"

func tracef(msg string, args ...interface{}) {
	s := fmt.Sprintf(msg, args...)
	if *trace {
		fmt.Printf("%s\n", s)
	}
}
