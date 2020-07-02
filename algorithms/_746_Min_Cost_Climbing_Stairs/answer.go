package _746_Min_Cost_Climbing_Stairs

func minCostClimbingStairs(cost []int) int {
	return minCostClimbingStairsDynamic(cost)
}

func minCostClimbingStairsDynamic(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		return minCost(cost[0], cost[1])
	}
	costList := make([]int, len(cost))
	costList[0] = cost[0]
	costList[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		costList[i] = minCost(costList[i-2]+cost[i], costList[i-1]+cost[i])
	}
	return minCost(costList[len(cost)-1], costList[len(cost)-2])
}

func minCost(a, b int) int {
	if a > b {
		return b
	}
	return a
}
