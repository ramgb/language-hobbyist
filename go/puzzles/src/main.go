package main

import (
	"fmt"
	"puzzles/src/hanoi"
	"puzzles/src/knights"
	"puzzles/src/nqueens"
)

func main() {
	fmt.Println("Welcome to puzzles!!")
	fmt.Println("1. Towers of Hanoi Solver")
	fmt.Println("2. N-Queens Solver")
	fmt.Println("3. Knight's Tour Solver")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		hanoi.Start()
	case 2:
		nqueens.Start()
	case 3:
		knights.Start()
	default:
		fmt.Println("Invalid choice")
	}
}
