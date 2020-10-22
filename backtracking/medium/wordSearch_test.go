package medium

import "testing"

func TestWordSearch(t *testing.T){
	board := [][]byte{
		[]byte{'A','B','C','E'},
		[]byte{'S','F','C','S'},
		[]byte{'A','D','E','E'},
	}
	if !exist(board, "ABCCED"){
		t.Fail()
	}
	if !exist(board, "SEE"){
		t.Fail()
	}
	if exist(board, "ABCB"){
		t.Fail()
	}
}
