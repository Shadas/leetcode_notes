package _127_Word_Ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	return ladderLengthBFS(beginWord, endWord, wordList)
}

func ladderLengthBFS(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	startIdx := len(wordList) - 1
	// 寻找结束单次的位置
	endIdx := -1
	for idx, word := range wordList {
		if word == endWord {
			endIdx = idx
			break
		}
	}
	if endIdx == -1 {
		return 0
	}
	// 构建邻接表
	adj := make(map[string][]int)
	for i, word := range wordList {
		for j := range word {
			source := word[:j] + "." + word[j+1:]
			adj[source] = append(adj[source], i)
		}
	}
	//
	var q []int
	q = append(q, startIdx)
	// distance[i] 表示从 startIndex 转换到 i 时的序列长度，
	// 初始化为 0 ，表示无法转换
	distance := make([]int, len(wordList))
	// 开始单词本身的无需任何转换就能得到
	distance[startIdx] = 1

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == endIdx {
			return distance[cur]
		}
		for j := range wordList[cur] {
			source := wordList[cur][:j] + "." + wordList[cur][j+1:]
			for _, next := range adj[source] {
				if distance[next] == 0 { // 该节点还没遍历过
					distance[next] = distance[cur] + 1
					q = append(q, next)
				}
			}
		}
	}
	// 遍历完未找到，则返回0
	return 0
}
