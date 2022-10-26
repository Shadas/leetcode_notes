package _862_Shortest_Subarray_with_Sum_at_Least_K

type mem struct {
	idx int
	sum int
}

func shortestSubarray(nums []int, k int) int {
	sum := 0
	queue := []mem{{idx: -1, sum: sum}}
	res := len(nums) + 1
	for idx, num := range nums {
		sum += num
		for len(queue) > 0 && sum-queue[0].sum >= k {
			res = min(res, idx-queue[0].idx)
			queue = queue[1:]
		}
		for len(queue) > 0 && sum <= queue[len(queue)-1].sum {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, mem{idx: idx, sum: sum})
	}

	if res == len(nums)+1 {
		res = -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
