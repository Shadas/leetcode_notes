package _128_Word_Ladder_2

import "fmt"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	return findLaddersBFS(beginWord, endWord, wordList)
}

func findLaddersBFS(beginWord string, endWord string, wordList []string) [][]string {
	wordList = append(wordList, beginWord)
	startIdx := len(wordList) - 1
	//
	endIdx := -1
	for idx, x := range wordList {
		if x == endWord {
			endIdx = idx
			break
		}
	}
	if endIdx == -1 {
		return nil
	}
	// 构建邻接表
	adj := make(map[string][]int)
	for i, word := range wordList {
		for j := range word {
			source := calcSource(word, j)
			adj[source] = append(adj[source], i)
		}
	}
	//
	distance := make([]int, len(wordList))
	distance[startIdx] = 1
	q := []int{startIdx}
	pres := make([][]int, len(wordList))

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for j := range wordList[cur] {
			source := calcSource(wordList[cur], j)
			for _, next := range adj[source] {
				if distance[next] == 0 {
					distance[next] = distance[cur] + 1
					pres[next] = append(pres[next], cur)
					q = append(q, next)
				} else if distance[next] == distance[cur]+1 {
					pres[next] = append(pres[next], cur)
				}
			}
		}
	}
	if distance[endIdx] == 0 {
		return nil
	}
	fmt.Println(pres)

	return nil
}

func calcSource(str string, idx int) string {
	return str[:idx] + "." + str[idx+1:]
}
