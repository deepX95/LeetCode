package simple

// waysToStep 面试题 08.01. 三步问题
func waysToStep(n int) int {
	// dp
	//if n <= 2 {
	//	return n
	//}
	//dp := make([]int, n+1)
	//dp[1], dp[2], dp[3] = 1, 2, 4
	//for i := 4; i <= n; i++ {
	//	dp[i] = (dp[i-1] + dp[i-2] + dp[i-3])%1000000007
	//}
	//return dp[n]

	// 滚动数组
	if n <= 2 {
		return n
	}
	a, c, d, e := 4, 0, 1, 2
	for i := 4; i <= n; i++ {
		c, d, e = d, e, a
		a = (c + d + e) % 1000000007

	}
	return a
}
