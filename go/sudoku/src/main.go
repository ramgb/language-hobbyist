package main

import (
	"fmt"
	"sudokusolver/src/sudoku"
	"sudokusolver/src/sudoku/solver"
)

func main() {
	sudoku := sudoku.NewSudoku("data/sudoku/1.txt")
	sudoku.PrintBoard()

	solverInstance := solver.NewSolver(sudoku, solver.BruteForceSolverType)

	time, guesses := solverInstance.Solve()
	fmt.Printf("%f seconds", time)
	fmt.Println()

	sudoku.SetSolvedBoard(solver.GetSolvedBoard(guesses))
	sudoku.PrintSolvedBoard()
}
