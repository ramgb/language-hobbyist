package main

import (
	"fmt"
	"sudokusolver/src/sudoku"
	"sudokusolver/src/sudoku/solver"
)

func main() {
	sudoku := sudoku.NewSudoku("data/sudoku/sample.txt")
	sudoku.PrintBoard()

	solver := solver.NewSimpleSolver(sudoku)

	fmt.Printf("%f seconds", solver.Solve())
	fmt.Println()

	sudoku.PrintSolvedBoard()
}
