package _875_Koko_Eating_Bananas

func minEatingSpeed(piles []int, H int) int {
	var (
		min, max, mid int
	)
	min = 1
	max = maxn(piles)
	if min == max {
		return min
	}
	for min < max {
		mid = (min + max) / 2
		if eatOver(piles, mid, H) {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return min
}

func eatOver(piles []int, k, H int) bool {
	total := 0
	for _, pile := range piles {
		if pile <= k {
			total += 1
		} else {
			total += pile/k + 1
		}
	}
	if total <= H {
		return true
	} else {
		return false
	}
}

func maxn(nums []int) (max int) {
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return
}
