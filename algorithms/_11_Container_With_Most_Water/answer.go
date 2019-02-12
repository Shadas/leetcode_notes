package _11_Container_With_Most_Water

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	container := 0
	for l < r {
		h := 0
		if height[l] > height[r] {
			h = height[r]
		} else {
			h = height[l]
		}
		if container < (r-l)*h {
			container = (r - l) * h
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}

	}
	return container
}
