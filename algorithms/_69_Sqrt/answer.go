package _69_Sqrt

func mySqrt(x int) int {
	var ret int
	for {
		if ret*ret > x {
			break
		}
		ret++
	}
	return ret - 1
}
