package _745_Prefix_and_Suffix_Search

type Node struct {
	next   [26]*Node
	weight map[string]int
}

type WordFilter struct {
	pr *Node
	sr *Node
}

func Constructor(words []string) WordFilter {
	pr := &Node{}
	sr := &Node{}
	for idx, word := range words {
		addWord(pr, word, idx)
		addWord(sr, reverseStr(word), idx)
	}
	return WordFilter{pr: pr, sr: sr}
}

func reverseStr(str string) string {
	ns := ""
	for _, b := range str {
		ns = string(b) + ns
	}
	return ns
}

func addWord(root *Node, word string, weight int) {
	var (
		cur = root
		i   int
	)
	for i < len(word) {
		b := word[i]
		idx := b - 'a'
		aimNode := cur.next[idx]
		if aimNode == nil {
			aimNode = &Node{weight: make(map[string]int)}
		}
		//aimNode.weight = append(aimNode.weight, weight)
		//tmpWord := word[:i+1]
		aimNode.weight[word] = weight
		cur.next[idx] = aimNode
		cur = aimNode
		i++
	}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	var rps, rss []int
	rps = searchStrFix(this.pr, prefix)
	rss = searchStrFix(this.sr, reverseStr(suffix))
	if len(rps) == 0 || len(rss) == 0 {
		return -1
	}
	var max int
	for _, rp := range rps {
		for _, rs := range rss {
			if rp == rs {
				if rs > max {
					max = rs
				}
			}
		}
	}
	return max
}

func searchStrFix(root *Node, word string) []int {
	var (
		cur = root
		i   int
	)
	for i >= 0 && i < len(word) {
		b := word[i]
		idx := b - 'a'
		aimNode := cur.next[idx]
		if aimNode == nil {
			return []int{}
		}
		cur = aimNode
		i++
	}
	weights := cur.weight
	var ret []int
	for _, v := range weights {
		ret = append(ret, v)
	}
	return ret
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
