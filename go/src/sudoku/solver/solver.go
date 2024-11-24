package solver

import (
	"exercises/src/sudoku"
)

type solver interface {
	solve()
	// get performance metrics function
}

type SimpleSolver struct {
	sudoku *sudoku.Sudoku
}
