package solver

import (
	"exercises/src/sudoku"
	"time"
)

type Solver interface {
	Solve() float64
}

type SimpleSolver struct {
	sudoku *sudoku.Sudoku
}

type ComplexSolver struct {
	sudoku *sudoku.Sudoku
}

func NewSimpleSolver(s *sudoku.Sudoku) *SimpleSolver {
	return &SimpleSolver{sudoku: s}
}

func (s *SimpleSolver) Solve() float64 {
	start := time.Now()
	return time.Since(start).Seconds()
}
