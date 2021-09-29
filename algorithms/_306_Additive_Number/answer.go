package _306_Additive_Number

func isAdditiveNumber(num string) bool {
	tmp := []int{}
	return dfs(num, 0, &tmp)
}

func dfs(num string, idx int, tmp *[]int) bool {
	if idx == len(num) && len(*tmp) > 2 {
		return true
	}
	cur := 0
	for i := idx; i < len(num); i++ {
		cur = cur*10 + int(num[i]-'0')
		if len(*tmp) >= 2 && cur-(*tmp)[len(*tmp)-1] > (*tmp)[len(*tmp)-2] {
			break
		}
		if len(*tmp) <= 1 || cur-(*tmp)[len(*tmp)-1] == (*tmp)[len(*tmp)-2] {
			*tmp = append(*tmp, cur)
			if dfs(num, i+1, tmp) {
				return true
			}
			*tmp = (*tmp)[0 : len(*tmp)-1]
		}
		if cur == 0 {
			break
		}
	}
	return false
}
