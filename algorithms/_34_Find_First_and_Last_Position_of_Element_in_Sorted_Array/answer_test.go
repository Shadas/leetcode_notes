package _34_Find_First_and_Last_Position_of_Element_in_Sorted_Array

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func TestFindFirst(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(FindFirst(nums, target, 0, len(nums)-1, -1))
}

func TestFindLast(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(FindLast(nums, target, 0, len(nums)-1, -1))
}
