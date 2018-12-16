package main

import "fmt"

func printShip() {
	switch game.period {
	case TOS:
		printEnterprise()
	case TNG:
		printEnterpriseD()
	}
}

func printEnterprise() {
	fmt.Printf("  __________________           __\n")
	fmt.Printf("  \\_________________|)____.---'--`---.____\n")
	fmt.Printf("                ||    \\----.________.----'\n")
	fmt.Printf("                ||     / /    `--'\n")
	fmt.Printf("              __||____/ /_\n")
	fmt.Printf("             |___ --====> \\\n")
	fmt.Printf("                 `--------'\n")
}

func printEnterpriseD() {
	fmt.Printf("                                     ____\n")
	fmt.Printf("                           __...---~'    `~~~----...__\n")
	fmt.Printf("                        _===============================\n")
	fmt.Printf("   ,----------------._/'      `---..._______...---'\n")
	fmt.Printf("   (_______________||_) . .  ,--'\n")
	fmt.Printf("       /    /.---'         `/\n")
	fmt.Printf("      '--------_- - - - - _/\n")
	fmt.Printf("                `--------'\n")
}
