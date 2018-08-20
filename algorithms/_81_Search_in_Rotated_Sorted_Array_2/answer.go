package _81_Search_in_Rotated_Sorted_Array_2

func search(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}
