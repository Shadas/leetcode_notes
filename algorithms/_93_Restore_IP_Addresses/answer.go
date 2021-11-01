package _93_Restore_IP_Addresses

import "strconv"

func restoreIpAddresses(s string) []string {
	var res []string
	find(4, s, "", &res)
	return res
}

// 递归回溯
func find(level int, str, tmp string, res *[]string) {
	if level == 0 {
		if len(str) == 0 {
			*res = append(*res, tmp[1:])
			return
		}
	} else {
		for i := 1; i <= 3; i++ { // 每段长度1-3位
			if len(str) < i {
				break
			}
			ss := str[0:i]
			if !isValid(ss) {
				break
			}
			find(level-1, str[i:], tmp+"."+ss, res)
		}
	}
}

func isValid(s string) bool {
	if len(s) == 0 || len(s) > 3 || (len(s) > 1 && s[0] == '0') {
		return false
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if num < 0 || num > 255 {
		return false
	}
	return true
}
