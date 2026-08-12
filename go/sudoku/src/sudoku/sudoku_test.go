package sudoku

import (
	"strings"
	"testing"
)

func TestNewSudoku_Success(t *testing.T) {
	// Test with newlines and spaces to ensure they are cleaned properly
	content := "________91________________________________________________5______________________"

	s, err := NewSudoku(content)
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
			name:    "too short",
			content: `________91________`,
			wantErr: "invalid board: expected 81 characters, got 18",
		},
		{
			name: "too long",
			content: strings.Join([]string{
				"________9",
				"1________",
				"_________",
				"_________",
				"_________",
				"_________",
				"____5____",
				"_________",
				"_________",
				"1234", // extra 4 chars
			}, ""),
			wantErr: "invalid board: expected 81 characters, got 85",
		},
		{
			name: "invalid character",
			content: strings.Join([]string{
				"_______X9",
				"1________",
				"_________",
				"_________",
				"_________",
				"_________",
				"____5____",
				"_________",
				"_________",
			}, ""),
			wantErr: "character 'X' at position 8 is not valid",
		},
		{
			name: "duplicate in row",
			content: strings.Join([]string{
				"9_______9",
				"1________",
				"_________",
				"_________",
				"_________",
				"_________",
				"____5____",
				"_________",
				"_________",
			}, ""),
			wantErr: "duplicate value 9 found in row 1",
		},
		{
			name: "duplicate in col",
			content: strings.Join([]string{
				"9________",
				"9________",
				"_________",
				"_________",
				"_________",
				"_________",
				"_________",
				"_________",
				"_________",
			}, ""),
			wantErr: "duplicate value 9 found in column 1",
		},
		{
			name: "duplicate in 3x3 box",
			content: strings.Join([]string{
				"_9_______",
				"__9______",
				"_________",
				"_________",
				"_________",
				"_________",
				"_________",
				"_________",
				"_________",
			}, ""),
			wantErr: "duplicate value 9 found in 3x3 subgrid starting at cell (1,1)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewSudoku(tc.content)
			if err == nil {
				t.Fatal("expected error, got nil")
			}

			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("expected error containing %q, got %q", tc.wantErr, err.Error())
			}
		})
	}
}
