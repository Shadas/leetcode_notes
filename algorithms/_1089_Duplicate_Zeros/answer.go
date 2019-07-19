package _1089_Duplicate_Zeros

func duplicateZeros(arr []int) {
	var (
		tmp = []int{}
	)
	for _, i := range arr {
		tmp = append(tmp, i)
		if i == 0 {
			tmp = append(tmp, 0)
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = tmp[i]
	}
}
