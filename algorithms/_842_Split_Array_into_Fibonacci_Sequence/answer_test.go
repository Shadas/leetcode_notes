package _842_Split_Array_into_Fibonacci_Sequence

import (
	"testing"

	"github.com/gotoprosperity/algorithm/util"
)

func TestSplitIntoFibonacci(t *testing.T) {
	if ret := splitIntoFibonacci("123456579"); !util.IsIntArrayEqual(ret, []int{123, 456, 579}) {
		t.Errorf("wrong ret with %v", ret)
	}
	if ret := splitIntoFibonacci("0123"); !util.IsIntArrayEqual(ret, []int{}) {
		t.Errorf("wrong ret with %v", ret)
	}
	if ret := splitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"); !util.IsIntArrayEqual(ret, []int{}) {
		t.Errorf("wrong ret with %v", ret)
	}
}
