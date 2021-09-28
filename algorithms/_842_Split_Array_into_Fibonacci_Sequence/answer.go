package _842_Split_Array_into_Fibonacci_Sequence

import "math"

func splitIntoFibonacci(num string) []int {
	ret := []int{}
	dfs(num, 0, &ret)
	return ret
}

func dfs(num string, idx int, ret *[]int) bool {
	if idx == len(num) {
		return len(*ret) > 2 // 因为下面的第二个if条件，这儿需要>2
	}
	cur := 0
	for i := idx; i < len(num); i++ {
		cur = cur*10 + int(num[i]-'0')
		if cur > math.MaxInt32 { // 越界也返回错误
			break
		}
		if len(*ret) >= 2 && cur-(*ret)[len(*ret)-1] > (*ret)[len(*ret)-2] { // 此数大于之前的n-1、n-2的和，已经失败
			break
		}
		if len(*ret) <= 1 || cur-(*ret)[len(*ret)-1] == (*ret)[len(*ret)-2] { // 长度不够，不能够定义斐波那契规律 or 正好匹配斐波那契规律 => 找下一个数
			*ret = append(*ret, cur) // 追加结果
			if dfs(num, i+1, ret) {  // 如果符合，则返回
				return true
			}
			// 此时下一个不符合，说明此次追加的结果及后续不可用
			*ret = (*ret)[0 : len(*ret)-1] // 撤销追加
			if cur == 0 {                  // 处理为0开头的情形
				break
			}
		}
		// 此时 小于之前n-1，n-2的和，继续循环加数尝试
	}
	return false
}
