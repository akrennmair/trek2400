package main

import (
	"bytes"
	"fmt"
)

type cvntab struct {
	abrev     string
	full      string
	boolValue bool
	intValue  int
}

var Yntab = []cvntab{
	{abrev: "y", full: "yes", boolValue: true},
	{abrev: "n", full: "no", boolValue: false},
}

func getynpar(s string) bool {
	r := getcodpar(s, Yntab)
	return r.boolValue
}

func getcodpar(s string, tab []cvntab) cvntab {
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
				return t
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
