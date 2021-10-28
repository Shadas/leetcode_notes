package _79_Word_Search

import "testing"

func TestExist(t *testing.T) {
	var (
		board [][]byte
		word  string
	)
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word = "ABCCED"
	if b := exist(board, word); !b {
		t.Errorf("should be true")
	}
	word = "SEE"
	if b := exist(board, word); !b {
		t.Errorf("should be true")
	}

	board = [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}
	word = "AAB"
	if b := exist(board, word); !b {
		t.Errorf("should be true")
	}

	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	word = "ABCESEEEFS"
	if b := exist(board, word); !b {
		t.Errorf("should be true")
	}
}
