package _316_Remove_Duplicate_Letters

import "strings"

func removeDuplicateLetters(s string) string {
	count := make([]int, 26)
	for _, b := range s {
		count[b-'a'] += 1
	}
	stackCount := make([]int32, 26)
	stack := make([]int32, len(s))
	index := -1

	for _, b := range s {
		for index != -1 && stack[index] > b && count[stack[index]-'a'] > 0 && stackCount[b-'a'] == 0 {
			stackCount[stack[index]-'a'] -= 1
			index--
		}
		if stackCount[b-'a'] == 0 {
			index++
			stack[index] = b
			stackCount[b-'a'] += 1
		}
		count[b-'a'] -= 1
	}

	builder := strings.Builder{}
	for i := 0; i <= index; i++ {
		builder.WriteByte(byte(stack[i]))
	}
	return builder.String()
}
