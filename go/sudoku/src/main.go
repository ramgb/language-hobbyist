package main

import (
	"log"
	"sudokusolver/src/gui"
	"sudokusolver/src/sudoku"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	s, err := sudoku.NewSudoku(sudoku.Easy)
	if err != nil {
		log.Fatalf("Failed to initialize Sudoku: %v", err)
	}

	g := gui.NewGame(s)
	ebiten.SetWindowSize(540, 620)
	ebiten.SetWindowTitle("Sudoku Solver (Ebitengine)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("GUI error: %v", err)
	}
}
