package _52_N_Queens_2

func totalNQueens(n int) int {
	return totalNQueensWithDFS(n)
}

func totalNQueensWithDFS(n int) int {
	var (
		pos []int   // 描述放置结果，下标为行，值为列，按行遍历dfs
		ret [][]int // 结果集
	)
	dfs(n, pos, &ret)
	return len(ret)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs(n int, pos []int, ret *[][]int) {
	if len(pos) == n {
		newPos := make([]int, len(pos))
		copy(newPos, pos)
		*ret = append(*ret, newPos)
		return
	}
column_loop:
	for i := 0; i < n; i++ { // 尝试第i列往里放
		// 检查是否有同列
		for _, p := range pos {
			if i == p {
				continue column_loop
			}
		}
		// 判断是否有存在对角线
		existDiagonal := false
		for line, col := range pos {
			if abs(len(pos)-line) == abs(col-i) {
				existDiagonal = true
				break
			}
		}
		// 如果有，尝试下一个位置
		if existDiagonal {
			continue
		}
		// 可以安放
		pos = append(pos, i)
		// 尝试安放下一个
		dfs(n, pos, ret)
		// 回退，尝试其他可能
		pos = pos[:len(pos)-1]
	}
	return
}
