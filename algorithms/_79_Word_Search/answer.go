package _79_Word_Search

func exist(board [][]byte, word string) bool {
	var (
		m, n   int // 行、列
		record [][]bool
	)
	m = len(board)
	if m == 0 {
		return false
	}
	n = len(board[0])
	if n == 0 {
		return false
	}
	if len(word) > m*n || len(word) == 0 {
		return false
	}
	// 构建record
	record = initRecord(m, n)
	// 遍历找开始
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				b := search(board, record, i, j, word[1:])
				if b {
					return true
				}
				record[i][j] = false
			}
		}
	}
	return false
}

func initRecord(m, n int) [][]bool {
	record := make([][]bool, m)
	for i := 0; i < m; i++ {
		record[i] = make([]bool, n)
	}
	return record
}

func availablePos(i, j int, record [][]bool) [][]int {
	m, n := len(record), len(record[0])
	ret := [][]int{}
	// 左
	if j-1 >= 0 && !record[i][j-1] {
		ret = append(ret, []int{i, j - 1})
	}
	// 上
	if i-1 >= 0 && !record[i-1][j] {
		ret = append(ret, []int{i - 1, j})
	}
	// 右
	if j+1 <= n-1 && !record[i][j+1] {
		ret = append(ret, []int{i, j + 1})
	}
	// 下
	if i+1 <= m-1 && !record[i+1][j] {
		ret = append(ret, []int{i + 1, j})
	}
	return ret
}

func search(board [][]byte, record [][]bool, i, j int, word string) bool {
	record[i][j] = true
	if len(word) == 0 {
		return true
	}
	aps := availablePos(i, j, record)
	for _, ap := range aps {
		ii, jj := ap[0], ap[1]
		if board[ii][jj] == word[0] {
			b := search(board, record, ii, jj, word[1:])
			if b {
				return true
			}
			record[ii][jj] = false
		}
	}
	return false
}
