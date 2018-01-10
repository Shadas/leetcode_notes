package _14_Longest_Common_Prefix

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix string = strs[0]
	var i = 1
	for {
		if i >= len(strs) {
			break
		}
		for {
			if idx := strings.Index(strs[i], prefix); idx != 0 {
				prefix = prefix[0 : len(prefix)-1]
			} else {
				break
			}
		}
		i++
	}
	return prefix
}
