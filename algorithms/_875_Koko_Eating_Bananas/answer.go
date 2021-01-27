package _875_Koko_Eating_Bananas

import "math"

func minEatingSpeed(piles []int, H int) int {
	var (
		min = 1
		max = math.MaxInt32
		mid int
	)
	for mid != (min+max)>>1 {
		mid = (min + max) >> 1
		if eatOver(piles, mid, H) {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return mid
}

func eatOver(piles []int, k, H int) bool {
	total := 0
	for _, pile := range piles {
		if pile <= k {
			total += 1
		} else {
			total += pile/k + 1
		}
		if total > H {
			return false
		}
	}
	return total <= H
}
