package _692_Top_K_Frequent_Words

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	return topKFrequentWithQS(words, k)
}

// 使用堆/pq做
func topKFrequentWithPQ(words []string, k int) []string {
	// todo: 思路比较明确，暂不实现
	return []string{}
}

// 使用快排来做
func topKFrequentWithQS(words []string, k int) []string {
	// 计数
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	// 放进数组
	s := make(X, len(m))
	idx := 0
	for k, v := range m {
		item := &Item{word: k, count: v}
		s[idx] = item
		idx++
	}
	// 排序
	sort.Sort(s)
	// 取前k个
	ret := make([]string, k)
	for i := 0; i < k; i++ {
		ret[i] = s[i].word
	}
	return ret
}

type Item struct {
	word  string
	count int
}

type X []*Item

func (x X) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x X) Len() int {
	return len(x)
}

func (x X) Less(i, j int) bool {
	return Less(x[i], x[j])
}

func Less(a, b *Item) bool {
	if a.count > b.count {
		return true
	}
	if a.count < b.count {
		return false
	}
	return a.word < b.word
}

func printL(l []*Item) {
	for _, x := range l {
		fmt.Printf("%s_%d,", x.word, x.count)
	}
	fmt.Println()
}

// 快排至k位，但这道题要求全有序，所以还是需要全排序
func SortUtilK(s []*Item, k int) {
	low, high := 0, len(s)
sortLoop:
	for low < high {
		pos := low
		i, j := low, high-1
		privot := s[pos]
		for i < j {
			for i <= j && Less(s[j], privot) {
				j--
			}
			if i <= j {
				s[j], s[pos] = s[pos], s[j]
				pos = j
			}
			for i <= j && Less(privot, s[i]) {
				i++
			}
			if i <= j {
				s[i], s[pos] = s[pos], s[i]
				pos = i
			}
		}
		if pos == k-1 {
			break sortLoop
		} else if pos < k-1 {
			low = pos + 1
		} else {
			high = pos
		}
	}
}
