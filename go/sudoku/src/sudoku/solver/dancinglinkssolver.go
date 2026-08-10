package solver

import (
	"sudokusolver/src/sudoku"
)

type DancingLinksSolver struct {
	sudoku *sudoku.Sudoku
}

func NewDancingLinksSolver(s *sudoku.Sudoku) *DancingLinksSolver {
	return &DancingLinksSolver{
		sudoku: s,
	}
}

// Solve the Sudoku
func (s *DancingLinksSolver) Solve() (float64, [9][9]int, error) {
	panic("Not Implemented")
}
