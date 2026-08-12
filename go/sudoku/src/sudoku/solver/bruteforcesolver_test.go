package solver

import (
	"strings"
	"sudokusolver/src/sudoku"
	"testing"
)

func TestBruteForceSolve_Success(t *testing.T) {
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
	var sb strings.Builder
	for _, r := range string(content) {
		if r == '\n' || r == '\r' || r == ' ' || r == '\t' {
			continue
		}
		sb.WriteRune(r)
	}

	s, err := sudoku.NewSudoku(sb.String())
	if err != nil {
		t.Fatalf("expected no error loading board, got %v", err)
	}

	solverInstance := NewBruteForceSolver(s)
	_, solvedBoard, err := solverInstance.Solve()
	if err != nil {
		t.Fatalf("expected puzzle to be solved, got error: %v", err)
	}

	// Verify that the solved board is valid
	var sbSolved strings.Builder
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if solvedBoard[i][j] == 0 {
				sbSolved.WriteByte('_')
			} else {
				sbSolved.WriteByte(byte('0' + solvedBoard[i][j]))
			}
		}
	}

	_, err = sudoku.NewSudoku(sbSolved.String())
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
	content := strings.Join([]string{
		"_12______",
		"356______",
		"478______",
		"_________",
		"9________",
		"_________",
		"_________",
		"_________",
		"_________",
	}, "")

	s, err := sudoku.NewSudoku(content)
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
	content := strings.Join([]string{
		"________9",
		"1________",
		"_________",
		"_________",
		"_________",
		"_________",
		"____5____",
		"_________",
		"_________",
	}, "")

	s, err := sudoku.NewSudoku(content)
	if err != nil {
		b.Fatalf("failed to load Sudoku: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solverInstance := NewBruteForceSolver(s)
		_, _, _ = solverInstance.Solve()
	}
}
