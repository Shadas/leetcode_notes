package _74_Search_a_2D_Matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 { // 如果横竖都不到1列，非法输入
		return false
	}
	m, n := len(matrix), len(matrix[0]) // 行、列
	l, r := 0, m*n-1
	for l <= r {
		mid := l + (r-l)/2
		i, j := idxToPos(mid, n)
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func idxToPos(idx, n int) (i, j int) {
	if idx < 0 {
		return -1, -1
	}
	i = idx / n
	j = idx % n
	return
}
