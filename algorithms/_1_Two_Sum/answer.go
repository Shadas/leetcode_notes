package _1_Two_Sum

func twoSum(nums []int, target int) []int {
	ret := []int{}
	for i1, n1 := range nums {
		for i2, n2 := range nums {
			if i1 != i2 && (n1+n2) == target {
				ret = []int{i1, i2}
				return ret
			}
		}
	}
	return ret
}
