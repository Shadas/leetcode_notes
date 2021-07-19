package _704_Binary_Search

func search(nums []int, target int) int {
	return searchRecursion(nums, 0, len(nums), target)
}

func searchRecursion(nums []int, start, end, target int) int {
	if start > end || start >= len(nums) || end < 0 {
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] > target {
		return searchRecursion(nums, start, mid-1, target)
	} else if nums[mid] < target {
		return searchRecursion(nums, mid+1, end, target)
	} else {
		return mid
	}
}
