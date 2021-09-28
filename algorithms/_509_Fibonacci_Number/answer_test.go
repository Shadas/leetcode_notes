package _509_Fibonacci_Number

import "testing"

func TestFib(t *testing.T) {
	if ret := fib(2); ret != 1 {
		t.Errorf("wrong ret with %d", ret)
	}
	if ret := fib(3); ret != 2 {
		t.Errorf("wrong ret with %d", ret)
	}
	if ret := fib(4); ret != 3 {
		t.Errorf("wrong ret with %d", ret)
	}
}
