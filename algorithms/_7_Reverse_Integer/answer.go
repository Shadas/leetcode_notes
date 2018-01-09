package _7_Reverse_Integer

import (
	"math"
)

func reverse(x int) int {
	ret := 0
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	for {
		if x == 0 {
			break
		}
		tail := x % 10
		newret := ret*10 + tail
		ret = newret
		x = x / 10
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return ret
}
