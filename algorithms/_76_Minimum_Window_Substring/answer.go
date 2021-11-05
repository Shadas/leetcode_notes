package _76_Minimum_Window_Substring

func minWindow(s string, t string) string {
	//return minWindowWithSlidingWindow(s, t)
	return minWindowWithSlidingWindowFast(s, t)
}

// 提升对比效率的滑动窗口
func minWindowWithSlidingWindowFast(s, t string) string {
	if len(s) < len(t) { // bad case
		return ""
	}
	var (
		tcMap       = make(map[int32]int)
		count       int // 作为匹配数的计算值，用于确认是否满足匹配，避免对于tcmap的遍历操作。触发条件为 count == len(t)
		left, right int

		ret    string       // 返回值
		minLen = len(s) + 1 // 用作计算是否为最小长度的变量。不用len(ret)的方式计算，这样不用预置ret的初始长度，避免找不到的情况下特判返回为空
	)
	// 初始化target字符串组成结构
	for _, tt := range t {
		tcMap[tt]++
	}
	// 开始遍历，以右边界为准，在每一个右边界的位置，收敛好当前右界下，所有左界情况
	for right = 0; right < len(s); right++ {
		cur := int32(s[right])
		if _, ok := tcMap[cur]; ok {
			tcMap[cur]--         // 此位置尝试-1
			if tcMap[cur] >= 0 { // 如果-1有效，说明匹配到了
				count++
			}
			for count == len(t) { // 如果成功匹配
				if right-left+1 < minLen { // 如果更小，更新
					minLen = right - left + 1
					ret = s[left : left+minLen]
				}
				curLeft := int32(s[left])
				if _, ok := tcMap[curLeft]; ok { // 如果是匹配的值，说明会影响
					tcMap[curLeft]++
					if tcMap[curLeft] > 0 { // 说明不满足条件了，拉齐count值
						count--
					}
				}
				left++ // 左边界尝试右移缩小窗口
			}
		}
	}
	return ret
}

// 滑动窗口思路没问题，判断子窗口是否满足条件，重复计算比较多，会超时，需要优化
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
	sm := make(map[int32]int)
	for _, ss := range s {
		n := sm[ss]
		n++
		sm[ss] = n
	}
	for _, tt := range t {
		n := sm[tt]
		n--
		if n < 0 {
			return false
		}
		sm[tt] = n
	}
	return true
}
