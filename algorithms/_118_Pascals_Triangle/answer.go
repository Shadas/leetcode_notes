package _118_Pascals_Triangle

func generate(numRows int) [][]int {
	var ret = [][]int{}
	if numRows <= 0 {
		return ret
	}
	for i := 1; i <= numRows; i++ {
		tmp := []int{1}
		if i != 1 {
			for j := 1; j < i-1; j++ {
				tmp = append(tmp, ret[i-2][j]+ret[i-2][j-1])
			}
			tmp = append(tmp, 1)
		}
		ret = append(ret, tmp)
	}
	return ret
}
