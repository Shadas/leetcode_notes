package _96_Unique_Binary_Search_Trees

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var (
		nums = make([]int, n+1)
	)
	nums[0], nums[1] = 0, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			nums[i] = nums[i] + nums[j]*nums[i-1-j]
		}
	}
	return 0
}
