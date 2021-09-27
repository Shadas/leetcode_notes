package _692_Top_K_Frequent_Words

import (
	"testing"

	"github.com/shadas/leetcode_notes/utils/array"
)

func TestTopKFrequent(t *testing.T) {
	var (
		words, ret []string
		k          int
	)
	words, k = []string{"i", "love", "leetcode", "i", "love", "coding"}, 2
	ret = topKFrequent(words, k)
	if !array.IsStrArrayEqual(ret, []string{"i", "love"}) {
		t.Errorf("wrong ret=%v", ret)
	}
	words, k = []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4
	ret = topKFrequent(words, k)
	if !array.IsStrArrayEqual(ret, []string{"the", "is", "sunny", "day"}) {
		t.Errorf("wrong ret=%v", ret)
	}
}

func TestLess(t *testing.T) {
	var a, b *Item
	a, b = &Item{word: "i", count: 2}, &Item{word: "love", count: 2}
	if !Less(a, b) {
		t.Error("should be true")
	}
	a, b = &Item{word: "i", count: 2}, &Item{word: "love", count: 1}
	if Less(a, b) {
		t.Error("should be false")
	}
}

func TestSortUtilK(t *testing.T) {
	var (
		s []*Item
		k int
	)
	s, k = []*Item{&Item{"the", 4}, &Item{"day", 1}, &Item{"is", 3}, &Item{"sunny", 2}}, 4
	SortUtilK(s, k)
	//printL(s)

}
