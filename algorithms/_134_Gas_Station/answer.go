package _134_Gas_Station

func canCompleteCircuit(gas []int, cost []int) int {
	var (
		total int
		cur   int
		ret   int
	)
	for idx, g := range gas {
		total += g - cost[idx]
		cur += g - cost[idx]
		if cur < 0 {
			cur = 0
			ret = idx + 1
		}
	}
	if total < 0 {
		return -1
	}
	return ret
}
