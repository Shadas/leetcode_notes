package _127_Word_Ladder

import "testing"

func TestLadderLengthBFS(t *testing.T) {
	var beginWord, endWord, wordList, ret = "hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5
	if r := ladderLengthBFS(beginWord, endWord, wordList); r != ret {
		t.Errorf("wrong ret with %d, not %d", r, ret)
	}
}
