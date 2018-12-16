package main

import "fmt"

func score() int {
	var (
		u int
		t int
		s int
		r float64
	)

	fmt.Printf("\n*** Your score:\n")
	u = game.enemiesKilled
	s = param.enemyPower / 4 * u
	t = s
	if t != 0 {
		fmt.Printf("%d %ss killed\t\t\t%6d\n", u, names.enemy, t)
	}
	r = now.date - param.date
	if r < 1.0 {
		r = 1.0
	}
	r = float64(game.enemiesKilled) / r
	t = int(400 * r)
	s += t
	if t != 0 {
		fmt.Printf("Kill rate %.2f %ss/stardate  \t%6d\n", names.enemy, r, t)
	}
	r = float64(now.enemies)
	r /= float64(game.enemiesKilled + 1)
	t = int(-400 * r)
	s += t
	if t != 0 {
		fmt.Printf("Penalty for %d %ss remaining\t%6d\n", names.enemy, now.enemies, t)
	}
	if move.endgame > 0 {
		u = game.skill
		t = 100 * u
		s += t
		fmt.Printf("Bonus for winning a %s game\t\t%d\n", skilltab[u-1].full, t)
	}
	if game.killed {
		s -= 500
		fmt.Printf("Penalty for getting killed\t\t  -500\n")
	}
	u = game.killb
	t = -100 * u
	s += t
	if t != 0 {
		fmt.Printf("%d starbases killed\t\t\t%6d\n", u, t)
	}
	u = game.helps
	t = -100 * u
	s += t
	if t != 0 {
		fmt.Printf("%d calls for help\t\t\t%6d\n", u, t)
	}
	u = game.kills
	t = -5 * u
	s += t
	if t != 0 {
		fmt.Printf("%d stars destroyed\t\t\t%6d\n", u, t)
	}
	u = game.killinhab
	t = -150 * u
	s += t
	if t != 0 {
		fmt.Printf("%d inhabited starsystems destroyed\t%6d\n", u, t)
	}
	if ship.ship != ENTERPRISE {
		s -= 200
		fmt.Printf("penalty for abandoning ship\t\t  -200\n")
	}
	u = game.captives
	t = 3 * u
	s += t
	if t != 0 {
		fmt.Printf("%d %ss captured\t\t\t%6d\n", u, names.enemy, t)
	}
	u = game.deaths
	t = -u
	s += t
	if t != 0 {
		fmt.Printf("%d casualties\t\t\t\t%6d\n", u, t)
	}
	fmt.Printf("\n***  TOTAL\t\t\t%14d\n", s)
	return s
}
