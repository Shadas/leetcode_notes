package _287_Find_the_Duplicate_Number

func findDuplicate(nums []int) int {
	var (
		slow = nums[0]
		fast = nums[nums[0]]
	)
	for {
		if slow == fast {
			break
		}
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for {
		if slow == fast {
			return fast
		}
		fast = nums[fast]
		slow = nums[slow]
	}
	return -1
}
