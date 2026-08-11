package main

import (
	"flag"
	"fmt"
	"log"
	"sudokusolver/src/gui"
	"sudokusolver/src/sudoku"
	"sudokusolver/src/sudoku/solver"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	guiFlag := flag.Bool("gui", false, "run Sudoku solver in 2D GUI mode")
	flag.Parse()

	s, err := sudoku.NewSudoku("data/sudoku/1.txt")
	if err != nil {
		log.Fatalf("Failed to initialize Sudoku: %v", err)
	}

	if *guiFlag {
		g := gui.NewGame(s)
		ebiten.SetWindowSize(540, 620)
		ebiten.SetWindowTitle("Sudoku Solver (Ebitengine)")
		if err := ebiten.RunGame(g); err != nil {
			log.Fatalf("GUI error: %v", err)
		}
	} else {
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
}
