package _658_Find_K_Closest_Element

func findClosestElements(arr []int, k int, x int) []int {
	return findClosestElementsFromHeadAndTail(arr, k, x)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// 堆/优先队列解法

// 两边遍历区间解法
func findClosestElementsFromHeadAndTail(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1
	for right-left+1 > k {
		if abs(arr[left]-x) > abs(arr[right]-x) {
			left++
		} else {
			right--
		}
	}
	return arr[left : right+1]
}

// 二分搜索解法
//func findClosestElementsBinarySearch(arr []int, k int, x int) []int {
//	left, right := 0, len(arr)-k
//	for left < right {
//		mid := left + (right-left)/2
//
//	}
//	return
//}
