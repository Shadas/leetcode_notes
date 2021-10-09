package _1201_Ugly_Number_3

import "math"

func nthUglyNumber(n int, a int, b int, c int) int {
	l, r := 1, math.MaxInt32
	for l < r {
		mid := l + (r-l)/2
		if cnt(mid, a, b, c) < n {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func floorInt(x, y int) int {
	return int(math.Floor(float64(x) / float64(y)))
}

func lcm2(x, y int) int {
	return floorInt(x*y, gcd(x, y))
}

func lcm3(x, y, z int) int {
	res := lcm2(x, y)
	return lcm2(z, res)
}

func cnt(k, x, y, z int) int {
	return floorInt(k, x) + floorInt(k, y) + floorInt(k, z) - floorInt(k, lcm2(x, y)) - floorInt(k, lcm2(x, z)) -
		floorInt(k, lcm2(y, z)) + floorInt(k, lcm3(x, y, z))
}
