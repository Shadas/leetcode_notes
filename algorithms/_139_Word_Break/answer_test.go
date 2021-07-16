package _139_Word_Break

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	var str string
	var wordDict []string
	str = "leetcode"
	wordDict = []string{"leet", "code"}
	fmt.Println(wordBreak(str, wordDict))
}
