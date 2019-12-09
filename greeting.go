package main

import (
	"fmt"
	"math/rand"
)

func printEnemyGreeting() {
	fmt.Printf("\n%s: incoming message from the %s ship, sir!\n\n", period.comms, period.enemy)

	switch game.period {
	case TOS:
		fmt.Printf("   \"%s!\"\n\n", randomKlingonProverb())

		fmt.Printf("          ____\n")
		fmt.Printf("         |____\\_____           _O__\n")
		fmt.Printf("    \\---------------\\_________======\n")
		fmt.Printf("     \\        ______/         \\____)\n")
		fmt.Printf(" ____/  /\\  |/_\n")
		fmt.Printf(" }_____________}\n")
	case TNG:
		fmt.Printf("   \"%s!\"\n\n", randomRomulanMessage())

		fmt.Printf("                    ____.-----.n.----._________.-.______\n")
		fmt.Printf("           ___,----'     \\|||||  _______________________`-._\n")
		fmt.Printf("       .--'    n                 \\------------------, ` \\___`.\n")
		fmt.Printf("  __,-' _____________             \\          `-===|  \\@\\      \\\n")
		fmt.Printf(" /-----'  /        __\\______.......|____        ==|   \\_) -----\\\n")
		fmt.Printf("|@ .-----|        |         |      |_  /         -\\     =======\\`\n")
		fmt.Printf("|  `-----|        |_________|......|___\\___________\\          __\\\n")
		fmt.Printf(" \\-----.__\\__________/             |_______.--------`._       `. \\\n")
		fmt.Printf("  `-.__                           / _/                 `--.__   \\|\n")
		fmt.Printf("       `-._u                    .'-'                         `-. |\n")
		fmt.Printf("            `---.___________.--'                                \\|\n")
	case ENTE:
		fmt.Println("We are the Borg. Lower your shields and surrender your ships.\nWe will add your biological and technological distinctiveness to our own.\nYour culture will adapt to service us. Resistance is futile.")
		fmt.Println(`
	___________
   /-/_"/-/_/-/|
  /"-/-_"/-_//||
 /__________/|/|
 |"|_'='-]:+|/||
 |-+-|.|_'-"||//
 |[".[:!+-'=|//
 |='!+|-:]|-|/
  ----------`)
	}

	fmt.Printf("\n")
}

var proverbs = []string{
	"Klingons are born to fight and conquer",
	"We are Klingons",
	"When threatened, fight",
	"We fight to enrich the spirit",
	"To survive, we must expand",
	"There is nothing shameful in falling before a superior enemy",
	"A warrior fights to the death",
	"A Klingon warrior is always prepared to fight",
	"A Klingon does not run away from his battles",
	"A fool and his head are soon parted",
	"Revenge is a dish best served cold",
}

func randomKlingonProverb() string {
	return proverbs[rand.Intn(len(proverbs))]
}

var romulanMessages = []string{
	"We're annexing this quadrant in the name of the Romulan Star Empire",
	"You will surrender as prisoners of war",
}

func randomRomulanMessage() string {
	return romulanMessages[rand.Intn(len(romulanMessages))]
}
