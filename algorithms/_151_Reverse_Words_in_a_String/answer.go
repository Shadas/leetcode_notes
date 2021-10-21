package _151_Reverse_Words_in_a_String

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	nwords := []string{}
	for _, word := range words {
		if word == "" {
			continue
		}
		nwords = append([]string{word}, nwords...)
	}
	return strings.Join(nwords, " ")
}
