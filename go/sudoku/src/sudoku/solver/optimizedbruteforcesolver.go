package solver

import (
	"sudokusolver/src/sudoku"
)

type OptimizedBruteForceSolver struct {
	sudoku *sudoku.Sudoku
}

func NewOptimizedBruteForceSolver(s *sudoku.Sudoku) *OptimizedBruteForceSolver {
	return &OptimizedBruteForceSolver{
		sudoku: s,
	}
}

// Solve the Sudoku
func (s *OptimizedBruteForceSolver) Solve() (float64, [9][9]int, error) {
	panic("Not Implemented")
}
