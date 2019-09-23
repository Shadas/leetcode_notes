package _303_Range_Sum_Query_Immutable

type NumArray struct {
	result []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{
		result: nums,
	}
	for i := 1; i < len(nums); i++ {
		na.result[i] = na.result[i-1] + nums[i]
	}
	return na
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > j {
		return -1
	}
	if i == 0 {
		return this.result[j]
	}
	return this.result[j] - this.result[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
