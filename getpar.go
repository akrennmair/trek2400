package main

import (
	"bytes"
	"fmt"
	"os"
)

type cvntab struct {
	abrev     string
	full      string
	boolValue bool
	intValue  int
	funcValue func(int)
}

var Yntab = []cvntab{
	{abrev: "y", full: "yes", boolValue: true},
	{abrev: "n", full: "no", boolValue: false},
}

func getynpar(s string) bool {
	r := getcodpar(s, Yntab)
	return r.boolValue
}

func getcodpar(s string, tab []cvntab) *cvntab {
	flag := false
	for {
		f := testnl()
		flag = flag || f
		if flag {
			fmt.Printf("%s: ", s)
		}
		if f {
			stdin.ReadByte() /* throw out the newline */
		}

		input, err := readToken()
		if err != nil {
			panic(err)
		}

		if input == "?" {
			c := 4
			for _, t := range tab {
				fmt.Printf("%14.14s", t.full)
				c--
				if c > 0 {
					continue
				}
				c = 4
				fmt.Println("")
			}
			if c != 4 {
				fmt.Println("")
			}
			continue
		}

		for _, t := range tab {
			if t.abrev == input || t.full == input {
				return &t
			}
		}

		fmt.Printf("invalid input; ? for valid inputs\n")
	}
}

func testnl() bool {
	for {
		c, _ := stdin.ReadByte()
		if c == '\n' {
			break
		}

		if (c >= '0' && c <= '9') || c == '.' || c == '!' ||
			(c >= 'A' && c <= 'Z') ||
			(c >= 'a' && c <= 'z') || c == '-' {
			stdin.UnreadByte()
			return false
		}
	}

	stdin.UnreadByte()
	return true
}

func readToken() (string, error) {
	var buf bytes.Buffer

	for {
		c, err := stdin.ReadByte()
		if err != nil {
			return "", err
		}
		if c == ' ' || c == '\t' || c == ';' || c == '\n' {
			stdin.UnreadByte()
			break
		}

		buf.WriteByte(c)
	}

	return buf.String(), nil
}

/**
 **     scan for newline
 **/
func skiptonl(c byte) {
	for c != '\n' {
		var err error
		c, err = stdin.ReadByte()
		if err != nil {
			return
		}
	}
	stdin.UnreadByte()
}

func getintpar(s string) int {
	for {
		if testnl() && s != "" {
			fmt.Printf("%s: ", s)
		}
		var n int
		i, err := fmt.Scanf("%d", &n)
		if i < 0 || err != nil {
			os.Exit(1)
		}
		if i > 0 && testterm() {
			return n
		}
		fmt.Printf("invalid input; please enter an integer\n")
		skiptonl(0)
	}
}

func getfltpar(s string) float64 {
	for {
		if testnl() && s != "" {
			fmt.Printf("%s: ", s)
		}
		var d float64
		i, err := fmt.Scan("%lf", &d)
		if i < 0 || err != nil {
			os.Exit(1)
		}
		if i > 0 && testterm() {
			return d
		}
		fmt.Printf("invalid input; please enter a double\n")
		skiptonl(0)
	}
}

func testterm() bool {
	c, _ := stdin.ReadByte()
	if c != 0 {
		return true
	}
	if c == '.' {
		return false
	}
	if c == '\n' || c == ';' {
		stdin.UnreadByte()
	}
	return true
}

func readdelim(d byte) bool {
	for {
		c, _ := stdin.ReadByte()
		if c == 0 {
			break
		}
		if c == d {
			return true
		}
		if c == ' ' {
			continue
		}
		stdin.UnreadByte()
		break
	}
	return false
}

func getstrpar(s string) string {
	if s != "" {
		fmt.Printf("%s: ", s)
	}
	skiptonl(0)
	stdin.ReadByte()
	answer, err := readToken()
	if err != nil {
		panic(err)
	}
	return answer
}
