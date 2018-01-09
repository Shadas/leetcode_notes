package _58_Length_of_Last_Word

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	strs := strings.Split(s, " ")
	return len(strs[len(strs)-1])
}
