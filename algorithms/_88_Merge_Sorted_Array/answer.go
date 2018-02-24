package _88_Merge_Sorted_Array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i > -1 && j > -1 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j > -1 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

//This is not the solution of the 88 question.
func mergeTwoArrayIntoRet(nums1 []int, m int, nums2 []int, n int) {
	var i, j int
	var ret []int
	for {
		if i == m || j == n {
			break
		}
		if nums1[i] < nums2[j] {
			ret = append(ret, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			ret = append(ret, nums2[j])
			j++
		} else {
			ret = append(ret, nums1[i], nums2[j])
			i++
			j++
		}
	}
	if i == m && j < n {
		ret = append(ret, nums2[j:]...)
	} else if j == n && i < m {
		ret = append(ret, nums1[i:]...)
	}
	nums1 = ret
}
