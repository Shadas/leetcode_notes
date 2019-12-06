package _167_Two_Sum_2_Input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	var (
		start = 0
		end   = len(numbers) - 1
	)
	for start < end {
		tmp := numbers[start] + numbers[end]
		if tmp < target {
			start++
		} else if tmp > target {
			end--
		} else {
			return []int{start + 1, end + 1}
		}
	}
	return nil
}
