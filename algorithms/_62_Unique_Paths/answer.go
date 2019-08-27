package _62_Unique_Paths

func uniquePaths(m int, n int) int {
	var res = []int{}
	for i := 0; i < n; i++ {
		res = append(res, 0)
	}
	res[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			res[j] += res[j-1]
		}
	}
	return res[n-1]
}
