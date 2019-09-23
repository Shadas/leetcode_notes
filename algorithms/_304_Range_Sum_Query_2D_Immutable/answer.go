package _304_Range_Sum_Query_2D_Immutable

type NumMatrix struct {
	ret [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nm := NumMatrix{
		ret: matrix,
	}
	for i := 0; i < len(matrix); i++ {
		sum := 0 // 一行一行计算
		for j := 0; j < len(matrix[i]); j++ {
			sum += matrix[i][j]
			if i != 0 { // 如果不是第一行, 则，需要加上上一行同列的值
				nm.ret[i][j] = sum + nm.ret[i-1][j]
			} else {
				nm.ret[i][j] = sum
			}
		}
	}
	return nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var (
		a, b, c, d int
	)
	a = this.ret[row2][col2]
	if row1 == 0 && col1 == 0 {
		b, c, d = 0, 0, 0
	} else if row1 == 0 && col1 != 0 {
		b, d = 0, 0
		c = this.ret[row2][col1-1]
	} else if row1 != 0 && col1 == 0 {
		c, d = 0, 0
		b = this.ret[row1-1][col2]
	} else {
		b = this.ret[row1-1][col2]
		c = this.ret[row2][col1-1]
		d = this.ret[row1-1][col1-1]
	}
	return a - b - c + d
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
