package _34_Find_First_and_Last_Position_of_Element_in_Sorted_Array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	head, tail := 0, len(nums)-1
	ret := []int{
		FindFirst(nums, target, head, tail, -1),
		FindLast(nums, target, head, tail, -1),
	}
	return ret
}

func FindFirst(nums []int, target, head, tail, lastPos int) (pos int) {
	if head > tail {
		return lastPos
	}
	mid := (head + tail) / 2
	if nums[mid] == target {
		tail = mid - 1
		lastPos = mid
	} else if nums[mid] > target {
		tail = mid - 1
	} else if nums[mid] < target {
		head = mid + 1
	}
	pos = FindFirst(nums, target, head, tail, lastPos)
	return
}

func FindLast(nums []int, target, head, tail, lastPos int) (pos int) {
	if head > tail {
		return lastPos
	}
	mid := (head + tail) / 2
	if nums[mid] == target {
		head = mid + 1
		lastPos = mid
	} else if nums[mid] > target {
		tail = mid - 1
	} else {
		head = mid + 1
	}
	pos = FindLast(nums, target, head, tail, lastPos)
	return
}
