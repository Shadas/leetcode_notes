package _1_Two_Sum

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	arr := []int{3, 2, 4}
	target := 6
	ret := twoSum(arr, target)
	want := []int{1, 2}
	if len(ret) != 2 {
		t.Errorf("Your answer is %v, it should be %v", ret, want)
	}
	if IntArrayIndex(ret, 1) == -1 || IntArrayIndex(ret, 2) == -1 {
		t.Errorf("Your answer is %v, it should be %v", ret, want)
	}

}

func IntArrayIndex(arr []int, ele int) int {
	ret := -1
	for i, value := range arr {
		if value == ele {
			ret = i
			break
		}
	}
	return ret
}
