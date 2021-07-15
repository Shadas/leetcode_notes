package _131_Palindrome_Partitioning

import (
	"fmt"
	"testing"
)

func TestIsPalindromic(t *testing.T) {
	if b := isPalindromic(""); !b {
		t.Error("'' should be palindromic")
	}
	if b := isPalindromic("aa"); !b {
		t.Error("aa should be palindromic")
	}
	if b := isPalindromic("aaa"); !b {
		t.Error("aaa should be palindromic")
	}
	if b := isPalindromic("acbaa"); b {
		t.Error("acbaa should not be palindromic")
	}
	if b := isPalindromic("aab"); b {
		t.Error("aab should not be palindromic")
	}
}

func TestPartition(t *testing.T) {
	var ret [][]string
	var str = "cbbbcc"
	ret = partition(str)
	fmt.Println(str, ret)

}
