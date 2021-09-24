package _658_Find_K_Closest_Element

func findClosestElements(arr []int, k int, x int) []int {
	return findClosestElementsFromHeadAndTail(arr, k, x)
	//return findClosestElementsBinarySearch(arr, k, x)
}

// 两边遍历区间解法
func findClosestElementsFromHeadAndTail(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1
	for right-left+1 > k {
		if x-arr[left] > arr[right]-x {
			left++
		} else {
			right--
		}
	}
	return arr[left : right+1]
}

// 二分搜索解法
func findClosestElementsBinarySearch(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k // 定义一个区间
	for left < right {
		mid := left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x { // 注意这里不能用绝对值
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}
