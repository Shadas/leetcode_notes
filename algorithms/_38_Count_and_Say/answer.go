package _38_Count_and_Say

import (
	"fmt"
)

func countAndSay(n int) string {
	str := "1"
	for i := 1; i < n; i++ {
		str = countStr(str)
	}
	return str
}

func countStr(str string) string {
	var nowStr string
	var nowCount int
	var ret string
	for _, c := range str {
		s := string(c)
		if string(s) == nowStr {
			nowCount++
		} else {
			if len(nowStr) != 0 {
				ret += fmt.Sprintf("%v%v", nowCount, string(nowStr))
				nowCount = 0
			}
			nowStr = s
			nowCount++
		}
	}
	ret += fmt.Sprintf("%v%v", nowCount, string(nowStr))
	return ret
}
