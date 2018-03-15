package _46_Permutations

func permute(nums []int) [][]int {
	ret := [][]int{}
	backtrack(&ret, &[]int{}, nums)
	return ret
}

func backtrack(list *[][]int, tmpList *[]int, nums []int) {
	if len(*tmpList) == len(nums) {
		l := []int{}
		for _, v := range *tmpList {
			l = append(l, v)
		}
		*list = append(*list, l)
	} else {
		for _, num := range nums {
			if intInList(num, *tmpList) {
				continue
			}
			*tmpList = append(*tmpList, num)
			backtrack(list, tmpList, nums)
			*tmpList = (*tmpList)[:len(*tmpList)-1]
		}
	}
}

func intInList(num int, list []int) bool {
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}
