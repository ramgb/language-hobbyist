package solver

import (
	"exercises/src/sudoku"
	"fmt"
	"time"
)

type solver interface {
	Solve() float64
	// Implement PrintMetrics() method
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
				updated, _ := updateValue(i, j, board[i][j], &guesses)
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

func updateValue(x int, y int, value int, guesses *[9][9]map[int]bool) (bool, [9][9]map[int]bool) {
	guesses[x][y] = map[int]bool{value: true}
	changed := [9][9]map[int]bool{}
	for i := 0; i < 9; i++ {
		if changed[x][i] == nil {
			changed[x][i] = map[int]bool{}
		}
		if changed[i][y] == nil {
			changed[i][y] = map[int]bool{}
		}
		if i != y {
			if guesses[x][i][value] {
				changed[x][i][value] = true
			}
			delete(guesses[x][i], value)
			if len(guesses[x][i]) == 0 {
				return false, changed
			}
		}
		if i != x {
			if guesses[i][y][value] {
				changed[i][y][value] = true
			}
			delete(guesses[i][y], value)
			if len(guesses[i][y]) == 0 {
				return false, changed
			}
		}
	}

	x0 := x / 3 * 3
	y0 := y / 3 * 3
	for i := x0; i < x0+3; i++ {
		for j := y0; j < y0+3; j++ {
			if changed[i][j] == nil {
				changed[i][j] = map[int]bool{}
			}
			if i != x && j != y {
				if guesses[i][j][value] {
					changed[i][j][value] = true
				}
				delete(guesses[i][j], value)
				if len(guesses[i][j]) == 0 {
					return false, changed
				}
			}
		}
	}
	return true, changed
}

func (s *SimpleSolver) undoUpdate(changed [9][9]map[int]bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for value := range changed[i][j] {
				s.guesses[i][j][value] = true
			}
		}
	}
}

// Print the guessed board
func (s *SimpleSolver) getSolvedBoard() [9][9]int {
	board := [9][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = getKey(s.guesses[i][j])
		}
	}
	return board
}

func (s *SimpleSolver) validateSolution() bool {
	board := s.getSolvedBoard()
	sumRow := [9]int{}
	sumCol := [9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sumRow[i] += board[i][j]
			sumCol[i] += board[j][i]
		}
	}

	for i := 0; i < 9; i++ {
		if sumRow[i] != 45 || sumCol[i] != 45 {
			return false
		}
	}
	return true
}

func getKey(m map[int]bool) int {
	for key := range m {
		return key
	}
	return -1
}

func (s *SimpleSolver) solveInternal(x int, y int) bool {
	// boundary condition - move to next row
	if y == 9 {
		return s.solveInternal(x+1, 0)
	}
	// boundary condition - end of board
	if x == 9 {
		return true
	}
	guesses := s.guesses[x][y]
	if len(guesses) > 1 {
		for guess := range guesses {
			updated, changed := updateValue(x, y, guess, &s.guesses)
			if updated {
				if s.solveInternal(x, y+1) {
					return true
				}
			}
			s.undoUpdate(changed)
		}
		s.guesses[x][y] = guesses
	} else if len(guesses) == 1 {
		return s.solveInternal(x, y+1)
	}

	return false
}

// Main function to solve the Sudoku
func (s *SimpleSolver) Solve() float64 {
	start := time.Now()

	solved := s.solveInternal(0, 0)
	fmt.Println("Solved: ", solved)
	fmt.Println("Valid: ", s.validateSolution())

	return time.Since(start).Seconds()
}
