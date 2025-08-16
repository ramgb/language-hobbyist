package knights

import (
	"fmt"
	"sort"
	"strconv"
)

type BestCell struct {
	Row   int
	Col   int
	Score int
}

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
	neighbors := getSortedNeighbors(board, row, col)

	for index := range neighbors {
		neighbor := neighbors[index]
		if neighbor.Score != -1 {
			jumpSuccess = jumpSuccess || knightjump(board, neighbor.Row, neighbor.Col, pos+1)
			if jumpSuccess {
				return true
			}
		}
	}
	board[row][col] = 0
	return false
}

func getSortedNeighbors(board [][]int, row, col int) []BestCell {
	neighbors := make([]BestCell, 8)
	modifiers := []int{1, 2}
	operators := []string{"+", "-"}
	index := 0
	for i := range modifiers {
		for j := range operators {
			for k := range operators {
				rowIndex := executeOperation(row, modifiers[i], operators[j])
				colIndex := executeOperation(col, modifiers[i^1], operators[k])
				score := getScore(board, rowIndex, colIndex)
				neighbors[index] = BestCell{Row: rowIndex, Col: colIndex, Score: score}
				index++
			}
		}
	}
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Score < neighbors[j].Score
	})
	return neighbors
}

func executeOperation(x int, y int, op string) int {
	if op == "+" {
		return x + y
	}
	if op == "-" {
		return x - y
	}
	return -1
}

func getScore(board [][]int, row, col int) int {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board) {
		return -1
	}
	score := 0
	boardSize := len(board)
	if row+1 < boardSize {
		if col+2 < boardSize {
			if board[row+1][col+2] == 0 {
				score++
			}
		}

		if col-2 >= 0 {
			if board[row+1][col-2] == 0 {
				score++
			}
		}
	}

	if row+2 < boardSize {
		if col+1 < boardSize {
			if board[row+2][col+1] == 0 {
				score++
			}
		}

		if col-1 >= 0 {
			if board[row+2][col-1] == 0 {
				score++
			}
		}
	}

	if row-1 >= 0 {
		if col+2 < boardSize {
			if board[row-1][col+2] == 0 {
				score++
			}
		}
		if col-2 >= 0 {
			if board[row-1][col-2] == 0 {
				score++
			}
		}
	}

	if row-2 >= 0 {
		if col+1 < boardSize {
			if board[row-2][col+1] == 0 {
				score++
			}
		}

		if col-1 >= 0 {
			if board[row-2][col-1] == 0 {
				score++
			}
		}
	}
	return score
}

func Start() {
	var boardSize int
	fmt.Println("Enter board size: ")
	fmt.Scanln(&boardSize)
	fmt.Println("Solving Knights Tour for " + strconv.Itoa(boardSize) + " X " + strconv.Itoa(boardSize))
	solve(boardSize)
}
