package _1019_Next_Greater_Node_In_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	// 放进数组
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	// 开始计算
	//return nextLargerNodesListAsc(list)
	return nextLargerNodesListDesc(list)
}

// 正序遍历处理
func nextLargerNodesListAsc(list []int) (ret []int) {
	ret = make([]int, len(list))
	var s []int
	for idx, val := range list {
		for {
			if len(s) == 0 {
				s = append(s, idx)
				break
			}
			top := list[s[len(s)-1]]
			if val > top {
				ret[s[len(s)-1]] = val
				s = s[:len(s)-1] // pop
			} else {
				s = append(s, idx)
				break
			}
		}
	}
	for _, idx := range s {
		ret[idx] = 0
	}
	return
}

// 倒序遍历处理
func nextLargerNodesListDesc(list []int) (ret []int) {
	ret = make([]int, len(list))
	var s []int
	for i := len(list) - 1; i >= 0; i-- {
		val := list[i]
		for {
			if len(s) == 0 {
				s = append(s, val)
				ret[i] = 0
				break
			}
			top := s[len(s)-1]
			if val < top {
				ret[i] = top
				s = append(s, val)
				break
			} else {
				s = s[:len(s)-1] // pop
			}
		}
	}
	return
}
