package _268_Missing_Number

func missingNumber(nums []int) int {
	// return missingNumberWithAdd(nums)
	return missingNumberWithXor(nums)
}

func missingNumberWithAdd(nums []int) int {
	var (
		rightSum int
		realSum  int
	)
	rightSum = len(nums) * (len(nums) + 1) / 2
	for _, n := range nums {
		realSum += n
	}
	return rightSum - realSum
}

func missingNumberWithXor(nums []int) int {
	var ret = len(nums)
	for i, n := range nums {
		ret ^= n
		ret ^= i
	}
	return ret
}
