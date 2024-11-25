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
					// throw an error
					return nil
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
	return true
}

func (s *SimpleSolver) solveInternal(x int, y int) bool {
	// return when hitting the final cell
	if y == 9 {
		return s.solveInternal(x+1, 0)
	} else if x == 9 {
		return true
	} else {
		// If there are no possible values for the cell, return since its invalid
		if len(s.guesses[x][y]) == 0 {
			return false
		}
		// if the cell has more than one possible value guess before moving to the next cell
		if len(s.guesses[x][y]) > 1 {
			guessesSnapshot := s.guesses
			for guess := range s.guesses[x][y] {
				updated := updateValue(x, y, guess, &s.guesses)
				if updated {
					solved := s.solveInternal(x, y+1)
					if solved {
						return true
					}
				}
				oldSnapshot := guessesSnapshot
				s.guesses = oldSnapshot
			}
		}
		return s.solveInternal(x, y+1)
	}
}

func (s *SimpleSolver) Solve() float64 {
	start := time.Now()
	solved := s.solveInternal(0, 0)

	if !solved {
		// throw an error
		return -1
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(len(s.guesses[i][j]))
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return time.Since(start).Seconds()
}
