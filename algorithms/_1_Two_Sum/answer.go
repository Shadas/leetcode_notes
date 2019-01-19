package _1_Two_Sum

func twoSum(nums []int, target int) []int {
	ret := []int{}
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if k1 >= k2 {
				continue
			}
			if v1+v2 == target {
				ret = append(ret, k1, k2)
				return ret
			}
		}
	}
	return ret
}
