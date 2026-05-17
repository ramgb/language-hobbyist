package solver

import "sudokusolver/src/sudoku"

type Solver interface {
	Solve() (float64, [9][9]map[int]bool)
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

// Print the guessed board
func GetSolvedBoard(guesses [9][9]map[int]bool) [9][9]int {
	board := [9][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = getKey(guesses[i][j])
		}
	}
	return board
}

func getKey(m map[int]bool) int {
	for key := range m {
		return key
	}
	return -1
}
