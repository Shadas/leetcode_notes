package _670_Maximum_Swap

import (
	"math"
)

func maximumSwap(num int) int {
	var (
		currMax    int
		maxPos     int
		pos1, pos2 int
		l          = []int{}
		ret        int
	)
	for i := 0; num != 0; i++ {
		n := num % 10
		l = append(l, n)
		if n > currMax {
			currMax = n
			maxPos = i
		}
		if n < currMax {
			pos1 = i
			pos2 = maxPos
		}

		num = num / 10
	}

	// 交换
	l[pos1], l[pos2] = l[pos2], l[pos1]
	// 求和
	for i, n := range l {
		ret += int(math.Pow(10, float64(i))) * n
	}
	return ret

}
