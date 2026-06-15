package nqueens

import (
	"fmt"
	"strconv"
)

func solve(boardSize int) {
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}
	queens := make([]int, boardSize)
	for i := range queens {
		queens[i] = -1
	}
	solveHelper(board, queens, 0)
}

func solveHelper(board [][]int, queens []int, row int) bool {
	if row >= len(board) {
		printBoard(queens, board)
		return true
	}

	hasValidSquare := false

	for col := range board[row] {
		if isValidSquare(board, row, col) {
			updateBoard(board, queens, row, col)
			hasValidSquare = hasValidSquare || solveHelper(board, queens, row+1)
			if hasValidSquare {
				break
			}
			revertBoard(queens, board, row, col)
		}
	}

	return hasValidSquare
}

func printBoard(queens []int, board [][]int) {
	for i := range board {
		for j := range board[i] {
			if queens[i] == j {
				fmt.Print(" Q ")
			} else {
				fmt.Print(" - ")
			}
		}
		fmt.Println()
	}
}

func isValidSquare(board [][]int, row, col int) bool {
	if board[row][col] != 0 {
		return false
	}
	boardSize := len(board)

	for i := 0; i < boardSize; i++ {
		if board[row][i] == -1 {
			return false
		}
		if board[i][col] == -1 {
			return false
		}
		if row-i >= 0 {
			if col-i >= 0 {
				if board[row-i][col-i] == -1 {
					return false
				}
			}

			if col+i < boardSize {
				if board[row-i][col+i] == -1 {
					return false
				}
			}
		}

		if row+i < boardSize {
			if col-i >= 0 {
				if board[row+i][col-i] == -1 {
					return false
				}
			}

			if col+i < boardSize {
				if board[row+i][col+i] == -1 {
					return false
				}
			}
		}

	}

	return true
}

func updateBoard(board [][]int, queens []int, row int, col int) {
	queens[row] = col
	// Implementation to check all things in the row/col/diagonal(s)

	boardSize := len(board)

	for i := 0; i < boardSize; i++ {
		board[row][i]++
		board[i][col]++

		if row-i >= 0 {
			if col-i >= 0 {
				board[row-i][col-i]++
			}

			if col+i < boardSize {
				board[row-i][col+i]++
			}
		}

		if row+i < boardSize {
			if col-i >= 0 {
				board[row+i][col-i]++
			}

			if col+i < boardSize {
				board[row+i][col+i]++
			}
		}

	}
	board[row][col] = -1
}

func revertBoard(queens []int, board [][]int, row int, col int) {
	queens[row] = -1

	boardSize := len(board)

	for i := 0; i < boardSize; i++ {
		board[row][i]--
		board[i][col]--

		if row-i >= 0 {
			if col-i >= 0 {
				board[row-i][col-i]--
			}

			if col+i < boardSize {
				board[row-i][col+i]--
			}
		}

		if row+i < boardSize {
			if col-i >= 0 {
				board[row+i][col-i]--
			}

			if col+i < boardSize {
				board[row+i][col+i]--
			}
		}

	}
	board[row][col] = 0
}

func Start() {
	var boardSize int
	fmt.Println("Enter board size: ")
	fmt.Scanln(&boardSize)
	fmt.Println("Solving nQueens for " + strconv.Itoa(boardSize) + " X " + strconv.Itoa(boardSize))
	solve(boardSize)
}
