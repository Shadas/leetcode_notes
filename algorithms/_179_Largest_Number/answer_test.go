package _179_Largest_Number

import "testing"

func TestLargestNumber(t *testing.T) {
	if ret := largestNumber([]int{10, 2}); ret != "210" {
		t.Errorf("wrong ret with %s", ret)
	}
	if ret := largestNumber([]int{3, 30, 34, 5, 9}); ret != "9534330" {
		t.Errorf("wrong ret with %s", ret)
	}
	if ret := largestNumber([]int{1}); ret != "1" {
		t.Errorf("wrong ret with %s", ret)
	}
	if ret := largestNumber([]int{10}); ret != "10" {
		t.Errorf("wrong ret with %s", ret)
	}
	if ret := largestNumber([]int{0, 0}); ret != "0" {
		t.Errorf("wrong ret with %s", ret)
	}
}
