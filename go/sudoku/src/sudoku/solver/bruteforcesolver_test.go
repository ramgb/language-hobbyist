package solver

import (
	"os"
	"path/filepath"
	"strings"
	"sudokusolver/src/sudoku"
	"testing"
)

func TestBruteForceSolve_Success(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "board.txt")
	content := []byte(`________9
1________
_________
_________
_________
_________
____5____
_________
_________
`)
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	s, err := sudoku.NewSudoku(filePath)
	if err != nil {
		t.Fatalf("expected no error loading board, got %v", err)
	}

	solverInstance := NewBruteForceSolver(s)
	_, solvedBoard, err := solverInstance.Solve()
	if err != nil {
		t.Fatalf("expected puzzle to be solved, got error: %v", err)
	}

	// Verify that the solved board is valid
	tempFile2 := filepath.Join(tempDir, "solved.txt")
	var sb strings.Builder
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if solvedBoard[i][j] == 0 {
				sb.WriteByte('_')
			} else {
				sb.WriteByte(byte('0' + solvedBoard[i][j]))
			}
		}
		sb.WriteByte('\n')
	}
	if err := os.WriteFile(tempFile2, []byte(sb.String()), 0644); err != nil {
		t.Fatalf("failed to write solved board: %v", err)
	}

	_, err = sudoku.NewSudoku(tempFile2)
	if err != nil {
		t.Fatalf("solved board is invalid: %v", err)
	}

	// Ensure no cells are empty
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if solvedBoard[i][j] < 1 || solvedBoard[i][j] > 9 {
				t.Errorf("invalid value %d at (%d, %d) in solved board", solvedBoard[i][j], i, j)
			}
		}
	}
}

func TestBruteForceSolve_Unsolvable(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "unsolvable.txt")
	// This board passes basic validation but has no possible values for cell (0,0)
	content := []byte(`_12______
356______
478______
_________
9________
_________
_________
_________
_________
`)
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	s, err := sudoku.NewSudoku(filePath)
	if err != nil {
		t.Fatalf("expected no error loading valid layout, got %v", err)
	}

	solverInstance := NewBruteForceSolver(s)
	_, _, err = solverInstance.Solve()
	if err == nil {
		t.Fatal("expected solver error for unsolvable board, got nil")
	}

	if !strings.Contains(err.Error(), "unsolvable") && !strings.Contains(err.Error(), "conflicts") {
		t.Errorf("expected unsolvable error message, got %q", err.Error())
	}
}

func BenchmarkBruteForceSolve(b *testing.B) {
	// Prepare a standard Sudoku puzzle
	tempDir := b.TempDir()
	filePath := filepath.Join(tempDir, "board.txt")
	content := []byte(`________9
1________
_________
_________
_________
_________
____5____
_________
_________
`)
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		b.Fatalf("failed to write temp file: %v", err)
	}

	s, err := sudoku.NewSudoku(filePath)
	if err != nil {
		b.Fatalf("failed to load Sudoku: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solverInstance := NewBruteForceSolver(s)
		_, _, _ = solverInstance.Solve()
	}
}
