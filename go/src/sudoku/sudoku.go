package sudoku

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type Sudoku struct {
	board      [9][9]int
	difficulty Difficulty
}

func (s *Sudoku) GetDifficulty() Difficulty {
	return s.difficulty
}

func (s *Sudoku) GetBoard() [9][9]int {
	return s.board
}

func estimateDifficulty(_ *[9][9]int) Difficulty {
	// For now, we'll just return Hard for all boards
	return Hard
}

func readBoardFromFile(inputFile string, board *[9][9]int) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file")
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	lineCtr := 0
	for scanner.Scan() {
		if lineCtr > 8 {
			log.Fatal("Too many lines in the file")
		}
		line := scanner.Text()
		if len(line) != 9 {
			log.Fatal("Line length is not 9")
		}
		for pos, char := range line {
			switch char {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				val := int(char - '0')
				board[lineCtr][pos] = val
			case '_':
				board[lineCtr][pos] = 0
			default:
				log.Fatal("Invalid character in line")
			}
		}
		lineCtr++
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func initBoard(inputFile string) *Sudoku {
	board := [9][9]int{}
	readBoardFromFile(inputFile, &board)
	CurrentDifficulty := estimateDifficulty(&board)
	return &Sudoku{board, CurrentDifficulty}
}

func (s *Sudoku) PrintBoard() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Print an underscore for empty cells
			if s.board[i][j] == 0 {
				fmt.Print("_")
			} else {
				fmt.Print(s.board[i][j])
			}
			// Add spaces for better formatting for all but the last column
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func NewSudoku(inputFile string) *Sudoku {
	return initBoard(inputFile)
}
