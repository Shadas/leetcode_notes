package _204_Count_Primes

import "testing"

func TestCountPrimes(t *testing.T) {
	if ret := countPrimes(2); ret != 0 {
		t.Errorf("ret of %d is %d rather than %d", 2, ret, 0)
	}
	if ret := countPrimes(10); ret != 4 {
		t.Errorf("ret of %d is %d rather than %d", 10, ret, 4)
	}

}
