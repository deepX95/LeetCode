package simple

import "math"

// maxProfit 121. 买卖股票的最佳时机
// fixme:https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/bao-li-mei-ju-dong-tai-gui-hua-chai-fen-si-xiang-b/
func maxProfit(prices []int) int {
	max,min:=0,math.MaxInt16
	for i:=range prices{
		max=maxV(max,prices[i]-min)
		min=minV(min,prices[i])
	}
	return max
}

func maxV(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func minV(a,b int)int{
	if a>b{
		return b
	}
	return a
}