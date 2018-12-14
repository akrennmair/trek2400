package main

import (
	"fmt"
	"math/rand"
)

func printEnemyGreeting() {
	fmt.Printf("\nUhura: incoming message from the Klingon ship, sir!\n\n")
	fmt.Printf("   \"%s!\"\n\n", randomKlingonProverb())

	fmt.Printf("          ____\n")
	fmt.Printf("         |____\\_____           _O__\n")
	fmt.Printf("    \\---------------\\_________======\n")
	fmt.Printf("     \\        ______/         \\____)\n")
	fmt.Printf(" ____/  /\\  |/_\n")
	fmt.Printf(" }_____________}\n")

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
