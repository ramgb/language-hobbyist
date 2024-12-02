package main

import (
	"exercises/src/sudoku"
	"exercises/src/sudoku/solver"
	"fmt"
)

func main() {
	sudoku := sudoku.NewSudoku("data/sudoku/sample.txt")
	sudoku.PrintBoard()

	solver := solver.NewSimpleSolver(sudoku)

	fmt.Printf("%f seconds", solver.Solve())
	fmt.Println()

	sudoku.PrintSolvedBoard()
}
