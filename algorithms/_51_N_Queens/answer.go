package _51_N_Queens

func solveNQueens(n int) [][]string {
	return solveNQueensWithDFS(n)
}

func solveNQueensWithDFS(n int) [][]string {
	var (
		pos    []int   // 描述放置结果，下标为行，值为列，按行遍历dfs
		tmp    [][]int // 中间结果
		result [][]string
	)
	dfs(n, pos, &tmp)
	result = transform(n, tmp)
	return result
}

func transform(n int, tmp [][]int) [][]string {
	var result [][]string
	for _, ret := range tmp {
		res := []string{}
		for _, t := range ret {
			str := ""
			for i := 0; i < n; i++ {
				if i == t {
					str += "Q"
				} else {
					str += "."
				}
			}
			res = append(res, str)
		}
		result = append(result, res)
	}
	return result
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
