package _121_Best_Time_to_Buy_and_Sell_Stock

import (
	"math"
)

func maxProfit(prices []int) int {
	return maxProfitOnce(prices)
}

func maxProfitOnce(prices []int) int {
	var max, maxSoFar int
	for i := 1; i < len(prices); i++ {
		max += prices[i] - prices[i-1]
		max = int(math.Max(0, float64(max)))
		maxSoFar = int(math.Max(float64(max), float64(maxSoFar)))
	}
	return maxSoFar
}

// can work, the time is not good, is refused by leetcode.
func maxProfitForce(prices []int) int {
	var max int
	for i, _ := range prices {
		for j := i; j < len(prices); j++ {
			x := prices[j] - prices[i]
			max = int(math.Max(float64(x), float64(max)))
		}
	}
	return max
}
