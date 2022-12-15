package _42_Trapping_Rain_Water

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	return trapStack(height)
}

func trapStack(height []int) int {
	s := []int{}
	sum := 0
	for idx, h := range height {
		for len(s) != 0 && height[s[len(s)-1]] < h {
			tmpIdx := s[len(s)-1]
			s = s[:len(s)-1]
			if len(s) != 0 {
				sum += (min(h, height[s[len(s)-1]]) - height[tmpIdx]) * (idx - s[len(s)-1] - 1)
			}
		}
		s = append(s, idx)
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
