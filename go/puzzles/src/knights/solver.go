package knights

import (
	"fmt"
	"strconv"
)

func solve(boardSize int) {
	board := make([][]bool, boardSize)
	for i := range board {
		board[i] = make([]bool, boardSize)
	}
}

func Start() {
	var boardSize int
	fmt.Println("Enter board size: ")
	fmt.Scanln(&boardSize)
	fmt.Println("Solving Knights Tour for " + strconv.Itoa(boardSize) + " X " + strconv.Itoa(boardSize))
	solve(boardSize)
}
