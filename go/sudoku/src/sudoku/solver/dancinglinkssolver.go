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

// Main function to solve the Sudoku
func (s *DancingLinksSolver) Solve() (float64, [9][9]map[int]bool) {
	panic("Not Implemented")
}
