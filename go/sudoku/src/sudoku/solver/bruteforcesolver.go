package solver

import (
	"fmt"
	"sudokusolver/src/sudoku"
	"time"
)

type BruteForceSolver struct {
	sudoku *sudoku.Sudoku
}

func NewBruteForceSolver(s *sudoku.Sudoku) *BruteForceSolver {
	return &BruteForceSolver{
		sudoku: s,
	}
}

func getOnlyValue(mask uint16) int {
	for v := 1; v <= 9; v++ {
		if (mask & (1 << v)) != 0 {
			return v
		}
	}
	return -1
}

func setAndPropagate(x, y int, val int, guesses *[9][9]uint16) bool {
	guesses[x][y] = 1 << val

	// Row and Column
	for i := 0; i < 9; i++ {
		if i != y {
			if (guesses[x][i] & (1 << val)) != 0 {
				guesses[x][i] &= ^(1 << val)
				if guesses[x][i] == 0 {
					return false
				}
			}
		}
		if i != x {
			if (guesses[i][y] & (1 << val)) != 0 {
				guesses[i][y] &= ^(1 << val)
				if guesses[i][y] == 0 {
					return false
				}
			}
		}
	}

	// 3x3 Box
	x0 := (x / 3) * 3
	y0 := (y / 3) * 3
	for i := x0; i < x0+3; i++ {
		for j := y0; j < y0+3; j++ {
			if i != x && j != y {
				if (guesses[i][j] & (1 << val)) != 0 {
					guesses[i][j] &= ^(1 << val)
					if guesses[i][j] == 0 {
						return false
					}
				}
			}
		}
	}
	return true
}

func (s *BruteForceSolver) solveInternal(x int, y int, guesses [9][9]uint16) (bool, [9][9]uint16) {
	// boundary condition - move to next row
	if y == 9 {
		return s.solveInternal(x+1, 0, guesses)
	}
	// boundary condition - end of board
	if x == 9 {
		return true, guesses
	}

	mask := guesses[x][y]
	if mask == 0 {
		return false, guesses
	}

	// Try all viable candidates
	for val := 1; val <= 9; val++ {
		if (mask & (1 << val)) != 0 {
			nextGuesses := guesses
			if setAndPropagate(x, y, val, &nextGuesses) {
				solved, result := s.solveInternal(x, y+1, nextGuesses)
				if solved {
					return true, result
				}
			}
		}
	}

	return false, guesses
}

// Solve the Sudoku and return execution time, solved board, or error.
func (s *BruteForceSolver) Solve() (float64, [9][9]int, error) {
	start := time.Now()

	// Initialize guesses: 0x3FE represents candidates 1..9 are possible for all cells.
	var guesses [9][9]uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			guesses[i][j] = 0x3FE
		}
	}

	// Apply initial board clues
	board := s.sudoku.GetBoard()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				if !setAndPropagate(i, j, board[i][j], &guesses) {
					return 0, [9][9]int{}, fmt.Errorf("initial board has constraint conflicts")
				}
			}
		}
	}

	solved, finalGuesses := s.solveInternal(0, 0, guesses)
	if !solved {
		return 0, [9][9]int{}, fmt.Errorf("sudoku puzzle is unsolvable")
	}

	// Convert finalGuesses to [9][9]int
	var solvedBoard [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			solvedBoard[i][j] = getOnlyValue(finalGuesses[i][j])
		}
	}

	return time.Since(start).Seconds(), solvedBoard, nil
}
