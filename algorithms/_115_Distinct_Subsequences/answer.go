package _115_Distinct_Subsequences

func numDistinct(s string, t string) int {
	sl, tl := len(s), len(t)
	if sl == 0 {
		return 0
	}
	var cache [][]int
	// init
	for i := 0; i <= tl; i++ {
		line := []int{}
		for j := 0; j <= sl; j++ {
			if i == 0 {
				line = append(line, 1)
			} else {
				line = append(line, 0)
			}
		}
		cache = append(cache, line)
	}
	//
	for i := 1; i <= tl; i++ {
		for j := i; j <= sl; j++ {
			if s[j-1] != t[i-1] {
				cache[i][j] = cache[i][j-1]
			} else {
				cache[i][j] = cache[i][j-1] + cache[i-1][j-1]
			}
		}
	}
	return cache[tl][sl]

}
