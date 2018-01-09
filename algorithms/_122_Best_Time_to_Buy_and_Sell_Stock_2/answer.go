package _122_Best_Time_to_Buy_and_Sell_Stock_2

func maxProfit(prices []int) int {
	var total int
	for i := 0; i < len(prices)-1; i++ {
		if val := prices[i+1] - prices[i]; val > 0 {
			total += val
		}
	}
	return total
}
