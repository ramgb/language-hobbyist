package main

import (
	"exercises/src/sudoku"
)

func main() {
	sudoku.NewSudoku("data/sudoku/sample.txt").PrintBoard()
}
