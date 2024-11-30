package solver

import (
	"exercises/src/sudoku"
	"fmt"
	"time"
)

type solver interface {
	Solve() float64
}

type SimpleSolver struct {
	sudoku  *sudoku.Sudoku
	guesses [9][9]map[int]bool
}

func NewSimpleSolver(s *sudoku.Sudoku) *SimpleSolver {

	guesses := [9][9]map[int]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			guesses[i][j] = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}
		}
	}

	board := s.GetBoard()

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				updated := updateValue(i, j, board[i][j], &guesses)
				if !updated {
					panic("Invalid board")
				}
			}
		}
	}

	return &SimpleSolver{
		sudoku:  s,
		guesses: guesses,
	}
}

func updateValue(x int, y int, value int, guesses *[9][9]map[int]bool) bool {
	guesses[x][y] = map[int]bool{value: true}
	for i := 0; i < 9; i++ {
		if i != y {
			delete(guesses[x][i], value)
			if len(guesses[x][i]) == 0 {
				return false
			}
		}
		if i != x {
			delete(guesses[i][y], value)
			if len(guesses[i][y]) == 0 {
				return false
			}
		}
	}

	x0 := x / 3 * 3
	y0 := y / 3 * 3
	for i := x0; i < x0+3; i++ {
		for j := y0; j < y0+3; j++ {
			if i != x && j != y {
				delete(guesses[i][j], value)
				if len(guesses[i][j]) == 0 {
					return false
				}
			}
		}
	}
	return true
}

// For Debugging - remove before productionizing
// Print the number of guesses for each cell
func (s *SimpleSolver) printGuesses() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(len(s.guesses[i][j]))
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Main function to solve the Sudoku
func (s *SimpleSolver) Solve() float64 {
	start := time.Now()

	// TODO: Implement the internal solve method
	s.printGuesses()
	return time.Since(start).Seconds()
}
