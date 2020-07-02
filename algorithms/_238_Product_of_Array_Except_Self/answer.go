package _238_Product_of_Array_Except_Self

func productExceptSelf(nums []int) []int {
	var (
		leftProduct = make([]int, len(nums))
		tmp         = 1
	)
	leftProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		tmp = tmp * nums[i+1]
		leftProduct[i] = tmp * leftProduct[i]
	}
	return leftProduct
}
