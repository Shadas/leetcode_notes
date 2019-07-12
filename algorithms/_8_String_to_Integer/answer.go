package _8_String_to_Integer

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	if strings.TrimSpace(str) == "" {
		return 0
	}
	var (
		idx  int
		ret  int64
		flag int64 = 1
	)
	// 去掉开头空字符
	for idx < len(str) && str[idx] == ' ' {
		idx++
	}
	if str[idx] == '-' {
		idx++
		flag = -1
	} else if str[idx] == '+' {
		idx++
	}
	for idx < len(str) && str[idx] >= '0' && str[idx] <= '9' {
		ret = ret*10 + int64(int(str[idx])-int('0'))
		if ret > math.MaxInt32 {
			if flag == -1 {
				ret = math.MinInt32
				flag = 1
			} else {
				ret = math.MaxInt32
			}
			break
		}
		idx++
	}
	ret *= flag
	return int(ret)
}
