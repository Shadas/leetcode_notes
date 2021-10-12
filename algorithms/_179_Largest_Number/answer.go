package _179_Largest_Number

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	s := make(StrNums, len(nums))
	for idx, num := range nums {
		s[idx] = strconv.FormatInt(int64(num), 10)
	}
	sort.Sort(s)
	str := ""
	for _, tmp := range s {
		str += tmp
	}
	str = strings.TrimLeft(str, "0")
	if str == "" {
		str = "0"
	}
	return str
}

type StrNums []string

func (s StrNums) Len() int {
	return len(s)
}

func (s StrNums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StrNums) Less(i, j int) bool {
	s1, s2 := s[i], s[j]
	if s1+s2 > s2+s1 {
		return true
	} else {
		return false
	}
}
