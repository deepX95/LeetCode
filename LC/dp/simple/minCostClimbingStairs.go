package simple

// MinCostClimbingStairs 746. 使用最小花费爬楼梯
func MinCostClimbingStairs(cost []int) int {
	// dp方程
	//n := len(cost)
	//dp := make([]int, n+1)
	//for i:=2;i<=n;i++{
	//	dp[i]= min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	//}
	//return dp[n]

	// 滚动数组
	n := len(cost)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
