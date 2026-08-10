package sudoku

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewSudoku_Success(t *testing.T) {
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

	s, err := NewSudoku(filePath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if s.GetDifficulty() != Hard {
		t.Errorf("expected difficulty to be Hard (3 clues), got %v", s.GetDifficulty())
	}

	board := s.GetBoard()
	if board[0][8] != 9 {
		t.Errorf("expected cell (0,8) to be 9, got %v", board[0][8])
	}
	if board[1][0] != 1 {
		t.Errorf("expected cell (1,0) to be 1, got %v", board[1][0])
	}
	if board[6][4] != 5 {
		t.Errorf("expected cell (6,4) to be 5, got %v", board[6][4])
	}
}

func TestNewSudoku_ValidationErrors(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr string
	}{
		{
			name: "too short",
			content: `________9
1________
`,
			wantErr: "too few lines in the file",
		},
		{
			name: "too long",
			content: `________9
1________
_________
_________
_________
_________
_________
_________
_________
_________
`,
			wantErr: "too many lines in the file",
		},
		{
			name: "invalid line length",
			content: `________
1________
_________
_________
_________
_________
_________
_________
_________
`,
			wantErr: "line 1 length is 8",
		},
		{
			name: "invalid character",
			content: `_______X9
1________
_________
_________
_________
_________
_________
_________
_________
`,
			wantErr: "character 'X' at line 1 position 8 is not valid",
		},
		{
			name: "duplicate in row",
			content: `9_______9
1________
_________
_________
_________
_________
_________
_________
_________
`,
			wantErr: "duplicate value 9 found in row 1",
		},
		{
			name: "duplicate in col",
			content: `9________
9________
_________
_________
_________
_________
_________
_________
_________
`,
			wantErr: "duplicate value 9 found in column 1",
		},
		{
			name: "duplicate in 3x3 box",
			content: `_9_______
__9______
_________
_________
_________
_________
_________
_________
_________
`,
			wantErr: "duplicate value 9 found in 3x3 subgrid starting at cell (1,1)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			filePath := filepath.Join(tempDir, "invalid_board.txt")
			if err := os.WriteFile(filePath, []byte(tc.content), 0644); err != nil {
				t.Fatalf("failed to write temp file: %v", err)
			}

			_, err := NewSudoku(filePath)
			if err == nil {
				t.Fatal("expected error, got nil")
			}

			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("expected error containing %q, got %q", tc.wantErr, err.Error())
			}
		})
	}
}
