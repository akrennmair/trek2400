package main

import "fmt"

func printShip() {
	switch game.period {
	case TOS:
		printEnterprise()
	case TNG:
		printEnterpriseD()
	case ENTE:
		printEnterpriseE()
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

func printEnterpriseE() {
	fmt.Printf("                                         __,----.____________________________\n")
	fmt.Printf("        ____.------------.____          /___\\ |>=============================|\n")
	fmt.Printf("  __.--'----------------------`--.__     `--.___.---,-----,-----------------'\n")
	fmt.Printf("======================================            .'   _.'\n")
	fmt.Printf("      `---------------.--------'     `----._____.'___.'____\n")
	fmt.Printf("                       `.========       >-----___________|_\\\n")
	fmt.Printf("                         `.                .-'\n")
	fmt.Printf("                           `--.______.----'\n")
}

func enemyShipIcon(c byte) byte {
	switch game.period {
	case TOS:
		return KLINGON
	case TNG:
		return ROMULAN
	case ENTE:
		return BORG
	default:
		return c
	}
}

func shipID() string {
	if ship.ship == QUEENE {
		return "NCC-1590" // 1590 is the year Faerie Queene was first published.
	}

	return period.shipid
}
