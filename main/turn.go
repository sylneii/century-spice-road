package main

import "fmt"

func turnManager(players int) func() int {
	turn := 0
	return func() int {
		if turn == players {
			turn = 1
		} else {
			turn += 1
		}
		return turn
	}
}

func turn() {
	nextturn := turnManager(4)
	fmt.Println(nextturn())
	fmt.Println(nextturn())
	fmt.Println(nextturn())
	fmt.Println(nextturn())
	fmt.Println(nextturn())
	fmt.Println(nextturn())
}
