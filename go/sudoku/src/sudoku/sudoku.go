package sudoku

import (
	"fmt"
	"os"
	"strings"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type Sudoku struct {
	board       [9][9]int
	solvedBoard [9][9]int
	difficulty  Difficulty
}

func (s *Sudoku) GetDifficulty() Difficulty {
	return s.difficulty
}

func (s *Sudoku) GetBoard() [9][9]int {
	return s.board
}

func (s *Sudoku) GetSolvedBoard() [9][9]int {
	return s.solvedBoard
}

func (s *Sudoku) SetSolvedBoard(board [9][9]int) {
	s.solvedBoard = board
}

func estimateDifficulty(board *[9][9]int) Difficulty {
	filledCells := 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				filledCells++
			}
		}
	}
	if filledCells < 26 {
		return Hard
	} else if filledCells < 32 {
		return Medium
	} else {
		return Easy
	}
}

func readBoardFromFile(inputFile string, board *[9][9]int) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Strip all newlines, carriage returns, spaces, and tabs
	var sb strings.Builder
	for _, r := range string(data) {
		if r == '\n' || r == '\r' || r == ' ' || r == '\t' {
			continue
		}
		sb.WriteRune(r)
	}
	cleaned := sb.String()

	if len(cleaned) != 81 {
		return fmt.Errorf("invalid board: expected 81 characters, got %d", len(cleaned))
	}

	for i, char := range cleaned {
		row := i / 9
		col := i % 9
		switch char {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val := int(char - '0')
			board[row][col] = val
		case '_', '.', '0':
			board[row][col] = 0
		default:
			return fmt.Errorf("invalid board: character %q at position %d is not valid (expected 1-9, '_', '.', or '0')", char, i+1)
		}
	}

	return nil
}

// Validate checks if the Sudoku board configuration violates any standard Sudoku rules.
func (s *Sudoku) Validate() error {
	// check rows
	for r := 0; r < 9; r++ {
		seen := 0
		for c := 0; c < 9; c++ {
			val := s.board[r][c]
			if val != 0 {
				bit := 1 << val
				if (seen & bit) != 0 {
					return fmt.Errorf("duplicate value %d found in row %d", val, r+1)
				}
				seen |= bit
			}
		}
	}
	// check columns
	for c := 0; c < 9; c++ {
		seen := 0
		for r := 0; r < 9; r++ {
			val := s.board[r][c]
			if val != 0 {
				bit := 1 << val
				if (seen & bit) != 0 {
					return fmt.Errorf("duplicate value %d found in column %d", val, c+1)
				}
				seen |= bit
			}
		}
	}
	// check 3x3 subgrids
	for box := 0; box < 9; box++ {
		seen := 0
		rowStart := (box / 3) * 3
		colStart := (box % 3) * 3
		for r := rowStart; r < rowStart+3; r++ {
			for c := colStart; c < colStart+3; c++ {
				val := s.board[r][c]
				if val != 0 {
					bit := 1 << val
					if (seen & bit) != 0 {
						return fmt.Errorf("duplicate value %d found in 3x3 subgrid starting at cell (%d,%d)", val, rowStart+1, colStart+1)
					}
					seen |= bit
				}
			}
		}
	}
	return nil
}

// IsValid returns true if the board is valid, false otherwise.
func (s *Sudoku) IsValid() bool {
	return s.Validate() == nil
}

func initBoard(inputFile string) (*Sudoku, error) {
	board := [9][9]int{}
	solvedBoard := [9][9]int{}
	if err := readBoardFromFile(inputFile, &board); err != nil {
		return nil, err
	}
	currentDifficulty := estimateDifficulty(&board)
	s := &Sudoku{board, solvedBoard, currentDifficulty}
	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("invalid initial board: %w", err)
	}
	return s, nil
}

func (s *Sudoku) PrintBoard() {
	s.internalPrintBoard(s.board)
}

func (s *Sudoku) PrintSolvedBoard() {
	s.internalPrintBoard(s.solvedBoard)
}

func (s *Sudoku) internalPrintBoard(anyBoard [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Print an underscore for empty cells
			if anyBoard[i][j] == 0 {
				fmt.Print("_")
			} else {
				fmt.Print(anyBoard[i][j])
			}
			// Add spaces for better formatting for all but the last column
			if j < 8 {
				fmt.Print("\t")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func NewSudoku(inputFile string) (*Sudoku, error) {
	return initBoard(inputFile)
}
