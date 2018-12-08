package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

const dumpFilename = "trek.dump"

var dumpTemplate = []interface{}{
	&ship, &now, &param, &etc, &game, &sect, &quad, &move, &eventList,
}

func dumpgame(v int) {
	f, err := os.OpenFile(dumpFilename, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("cannot dump: %s\n", err.Error())
		return
	}
	defer f.Close()

	enc := gob.NewEncoder(f)

	for _, data := range dumpTemplate {
		if err := enc.Encode(data); err != nil {
			fmt.Printf("cannot encode dump: %s\n", err.Error())
			return
		}
	}
}

func restartgame() bool {
	f, err := os.Open(dumpFilename)
	if err != nil {
		fmt.Printf("cannot restart: %s\n", err.Error())
		return false
	}
	defer f.Close()

	if !readdump(f) {
		fmt.Printf("cannot restart\n")
		return false
	}

	return true
}

func readdump(f io.Reader) bool {
	dec := gob.NewDecoder(f)

	for _, data := range dumpTemplate {
		if err := dec.Decode(data); err != nil {
			return false
		}
	}

	return true
}
