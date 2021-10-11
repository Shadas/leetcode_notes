package _300_Longest_Increasing_Subsequence

func lengthOfLIS(nums []int) int {
	//return lengthOfLISDP(nums)
	return lengthOfLISGreedy(nums)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp解法
func lengthOfLISDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		dp      = make([]int, len(nums))
		max int = 1
	)
	for idx, _ := range dp {
		dp[idx] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		max = maxInt(max, dp[i])
	}
	return max
}

// 贪心法构造具体数列解法
func lengthOfLISGreedy(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		lis = []int{nums[0]}
	)
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		head, tail := lis[0], lis[len(lis)-1]
		if val <= head {
			lis[0] = val
		} else if val > tail {
			lis = append(lis, val)
		} else {
			lis[findPos(lis, val)] = val
		}
	}
	return len(lis)
}

// 二分法找第一个>=val的位置
func findPos(lis []int, val int) int {
	l, r := 0, len(lis)-1
	for l < r {
		mid := l + (r-l)/2
		if lis[mid] >= val {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
