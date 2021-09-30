package _240_Search_a_2D_Matrix_2

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 { // 非法情形
		return false
	}
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	val := matrix[i][j]
	for i < m && j >= 0 {
		val = matrix[i][j]
		if val == target {
			return true
		} else if val > target {
			j--
		} else {
			i++
		}

	}
	return false
}
