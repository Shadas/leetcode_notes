package _704_Binary_Search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	var (
		nums   []int
		target int
	)
	nums = []int{-1, 0, 3, 5, 9, 12}
	target = 13
	fmt.Println(search(nums, target))

}
