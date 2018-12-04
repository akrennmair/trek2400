package main

import "fmt"

func out(dev int) {
	d := &devices[dev]
	fmt.Printf("%s reports %s ", d.person, d.name)
	if d.name[len(d.name)-1] == 's' {
		fmt.Printf("are")
	} else {
		fmt.Printf("is")
	}
	fmt.Printf(" damaged\n")
}
