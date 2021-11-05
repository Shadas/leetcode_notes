package _76_Minimum_Window_Substring

func minWindow(s string, t string) string {
	return minWindowWithSlidingWindow(s, t)
}

func minWindowWithSlidingWindow(s, t string) string {
	if len(s) < len(t) { // bad case
		return ""
	}
	var (
		ret  string = s + t
		p, q int    // [)
	)
	for q < len(s) || p < len(s) {
		if q-p < len(t) { // 长度不够，需要+1
			//fmt.Printf("not enough len with %s, p=%d, q=%d\n", s[p:q], p, q)
			if q == len(s) { // q不能加了，说明没了
				break
			} else { // 还能加
				q += 1
				continue
			}
		}
		tmp := s[p:q]
		//fmt.Printf("test str=%s, p=%d, q=%d\n", tmp, p, q)
		if isValid(tmp, t) { // 如果有，尝试缩短，看看是否有冗余空间
			if len(tmp) < len(ret) { // 看长短，看是否更新
				ret = tmp
			}
			p += 1
			continue
		} else { // 现在不存在，尝试扩展范围，类似长度不够的逻辑
			if q == len(s) { // q不能加了，说明没了
				break
			} else { // 还能加
				q += 1
				continue
			}
		}
	}
	if len(ret) > len(s) {
		return ""
	}

	return ret
}

func isValid(s, t string) bool {
	tm, sm := make(map[int32]int), make(map[int32]int)
	for _, ss := range s {
		n := sm[ss]
		n++
		sm[ss] = n
	}
	for _, tt := range t {
		n := tm[tt]
		n++
		tm[tt] = n
	}
	for tt, tn := range tm {
		if sm[tt] < tn {
			return false
		}
	}
	return true
}
