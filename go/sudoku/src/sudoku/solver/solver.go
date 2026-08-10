package solver

import "sudokusolver/src/sudoku"

type Solver interface {
	Solve() (float64, [9][9]int, error)
}

type solverType int

const (
	BruteForceSolverType          solverType = iota
	OptimizedBruteForceSolverType            = 1
	DancingLinksSolverType                   = 2
)

func NewSolver(s *sudoku.Sudoku, t solverType) Solver {
	switch t {
	case BruteForceSolverType:
		return NewBruteForceSolver(s)
	case OptimizedBruteForceSolverType:
		return NewOptimizedBruteForceSolver(s)
	case DancingLinksSolverType:
		return NewDancingLinksSolver(s)
	default:
		panic("Invalid solver type")
	}
}
