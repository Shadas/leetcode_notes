package _997_Find_the_Town_Judge

func findJudge(N int, trust [][]int) int {
	var (
		indegree  []int
		outdegree []int
	)
	for i := 0; i < N; i++ {
		indegree = append(indegree, 0)
		outdegree = append(outdegree, 0)
	}

	for _, t := range trust {
		pre, post := t[0], t[1]
		outdegree[pre-1]++
		indegree[post-1]++
	}

	for i := 0; i < N; i++ {
		if indegree[i] == N-1 && outdegree[i] == 0 {
			return i + 1
		}
	}
	return -1
}
