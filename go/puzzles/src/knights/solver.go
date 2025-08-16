package knights

import (
	"fmt"
	"strconv"
)

func solve(boardSize int) {
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}
	solved := knightjump(board, 0, 0, 1)
	if solved {
		printBoard(board)
	} else {
		fmt.Println("Cant solve")
	}
}

func printBoard(board [][]int) {
	for i := range board {
		for j := range board[i] {
			fmt.Print(strconv.Itoa(board[i][j]) + "\t")
		}
		fmt.Println()
	}
}

func knightjump(board [][]int, row, col, pos int) bool {
	boardSize := len(board)

	if pos > boardSize*boardSize {
		return true
	}
	if board[row][col] != 0 {
		return false
	}

	board[row][col] = pos

	if pos == boardSize*boardSize {
		return true
	}
	jumpSuccess := false

	// Implement warnsdorff's rule using a priority queue
	if row+1 < boardSize {
		if col+2 < boardSize {
			jumpSuccess = jumpSuccess || knightjump(board, row+1, col+2, pos+1)
		}
		if jumpSuccess {
			return true
		}

		if col-2 >= 0 {
			jumpSuccess = jumpSuccess || knightjump(board, row+1, col-2, pos+1)
		}

		if jumpSuccess {
			return true
		}
	}

	if row+2 < boardSize {

		if col+1 < boardSize {
			jumpSuccess = jumpSuccess || knightjump(board, row+2, col+1, pos+1)
		}

		if jumpSuccess {
			return true
		}

		if col-1 >= 0 {
			jumpSuccess = jumpSuccess || knightjump(board, row+2, col-1, pos+1)
		}

		if jumpSuccess {
			return true
		}
	}

	if row-1 >= 0 {
		if col+2 < boardSize {
			jumpSuccess = jumpSuccess || knightjump(board, row-1, col+2, pos+1)
		}

		if jumpSuccess {
			return true
		}
		if col-2 >= 0 {
			jumpSuccess = jumpSuccess || knightjump(board, row-1, col-2, pos+1)
		}

		if jumpSuccess {
			return true
		}
	}

	if row-2 >= 0 {
		if col+1 < boardSize {
			jumpSuccess = jumpSuccess || knightjump(board, row-2, col+1, pos+1)
		}

		if jumpSuccess {
			return true
		}

		if col-1 >= 0 {
			jumpSuccess = jumpSuccess || knightjump(board, row-2, col-1, pos+1)
		}

		if jumpSuccess {
			return true
		}
	}

	board[row][col] = 0
	return false
}

func Start() {
	var boardSize int
	fmt.Println("Enter board size: ")
	fmt.Scanln(&boardSize)
	fmt.Println("Solving Knights Tour for " + strconv.Itoa(boardSize) + " X " + strconv.Itoa(boardSize))
	solve(boardSize)
}
