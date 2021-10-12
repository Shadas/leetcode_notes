package _1260_Shift_2D_Grid

func shiftGrid(grid [][]int, k int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 { // bad case
		return nil
	}
	var (
		m, n int // 行、列数
		l    []int
		ret  [][]int
	)
	m = len(grid)
	n = len(grid[0])
	l = make([]int, m*n)
	ret = make([][]int, m)
	for i := 0; i < m; i++ {
		ret[i] = make([]int, n)
	}
	// 初始化
	for ln, line := range grid {
		for cn, x := range line {
			l[ln*n+cn] = x
		}
	}
	// 切换
	l = append(l[len(l)-k%len(l):], l[0:len(l)-k%len(l)]...)
	// 还原
	for idx, x := range l {
		l := idx / n
		c := idx - n*l
		ret[l][c] = x
	}
	return ret
}
