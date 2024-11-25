package _128_Word_Ladder_2

import "testing"

func TestFindLaddersBFS(t *testing.T) {
	findLaddersBFS("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
}
