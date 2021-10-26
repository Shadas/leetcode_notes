package _137_Single_Number_2

const Times = 3

func singleNumber(nums []int) int {
	//return singleNumberWithCalcu(nums)
	return singleNumberWithBit(nums)
}

// 位运算解法
func singleNumberWithBit(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		a = (a ^ num) & ^b
		b = (^a ^ num) ^ b
	}
	return a
}

// 算术求和解法，先去重，再*3，减去所有原来num和，即2倍的目标值。此处不予实现。
func singleNumberWithSum(nums []int) int {
	// ...
	return 0
}

// 计数解法
func singleNumberWithCalcu(nums []int) int {
	mark := []bool{}
	for i := 0; i < len(nums); i++ {
		mark = append(mark, false)
	}

	for i, n := range nums {
		if mark[i] == false {
			if i == len(nums)-1 {
				return n
			} else {
				tmp := nums[i+1:]
				times := 0
				for j, m := range tmp {
					if m == n {
						mark[i+j+1] = true
						times++
						if times == Times-1 {
							break
						}
					}
				}
				if times != Times-1 {
					return n
				}
			}
		} else {
			continue
		}
	}
	return 0
}
