package _121_Best_Time_to_Buy_and_Sell_Stock

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	if ret := maxProfit([]int{7, 1, 5, 3, 6, 4}); ret != 5 {
		t.Error("not 5 with", ret)
	}
}
