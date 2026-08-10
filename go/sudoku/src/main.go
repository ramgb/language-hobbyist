package main

import (
	"fmt"
	"log"
	"sudokusolver/src/sudoku"
	"sudokusolver/src/sudoku/solver"
)

func main() {
	s, err := sudoku.NewSudoku("data/sudoku/1.txt")
	if err != nil {
		log.Fatalf("Failed to initialize Sudoku: %v", err)
	}

	fmt.Println("Initial Board:")
	s.PrintBoard()

	solverInstance := solver.NewSolver(s, solver.BruteForceSolverType)

	duration, solvedBoard, err := solverInstance.Solve()
	if err != nil {
		log.Fatalf("Failed to solve Sudoku: %v", err)
	}

	fmt.Printf("Solved in %f seconds\n\n", duration)

	s.SetSolvedBoard(solvedBoard)
	fmt.Println("Solved Board:")
	s.PrintSolvedBoard()
}
