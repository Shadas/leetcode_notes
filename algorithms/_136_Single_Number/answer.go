package _136_Single_Number

func singleNumber(nums []int) int {
	//return singleNumberWithMap(nums)
	return singleNumberWithXOR(nums)
}

// 异或 计算处理
func singleNumberWithXOR(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}
	return ret
}

// 使用额外map处理
func singleNumberWithMap(nums []int) int {
	var m = make(map[int]int)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			delete(m, n)
		} else {
			m[n] = 0
		}
	}
	for r, _ := range m {
		return r
	}
	return 0
}
