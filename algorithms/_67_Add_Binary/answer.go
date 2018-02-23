package _67_Add_Binary

import (
	"fmt"
)

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	// let a >= b
	if la < lb {
		la, lb = lb, la
		a, b = b, a
	}
	ca, cb := []byte(a), []byte(b)
	var isAdd bool
	var ret []byte
	for i := lb - 1; i >= 0; i-- {
		var va, vb int
		var v string
		if string(ca[i+la-lb]) == "1" {
			va = 1
		}
		if string(cb[i]) == "1" {
			vb = 1
		}
		sum := va + vb
		if isAdd {
			sum++
		}
		switch sum {
		case 0:
			v = "0"
			isAdd = false
		case 1:
			v = "1"
			isAdd = false
		case 2:
			v = "0"
			isAdd = true
		case 3:
			v = "1"
			isAdd = true
		}
		ret = append([]byte(v), ret...)
	}
	for i := la - lb - 1; i >= 0; i-- {
		var v int
		if string(ca[i]) == "1" {
			v = 1
		}
		if isAdd {
			v = v + 1
			if v == 2 {
				isAdd = true
				v = 0
			} else {
				isAdd = false
			}
		}
		ret = append([]byte(fmt.Sprintf("%v", v)), ret...)
	}
	if isAdd {
		ret = append([]byte("1"), ret...)
	}
	return string(ret)
}
