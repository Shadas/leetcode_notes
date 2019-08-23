package _290_Word_Pattern

import "strings"

func wordPattern(pattern string, str string) bool {
	var pm = make(map[rune][]int)
	for idx, b := range pattern {
		idxs, _ := pm[b]
		idxs = append(idxs, idx)
		pm[b] = idxs
	}
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	var diffWord []string
	for _, idxs := range pm {
		tmp := strs[idxs[0]]
		for _, idx := range idxs {
			if strs[idx] != tmp {
				return false
			}
		}
		for _, dw := range diffWord {
			if tmp == dw {
				return false
			}
		}
		diffWord = append(diffWord, tmp)
	}

	return true
}
