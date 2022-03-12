package simple

// leastMinutes LCS 01. 下载插件
func leastMinutes(n int) int {
	// dp[i] 表示下载 i 个插件需要的最少分钟数
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		//dp[i] = dp[(i + 1) / 2] + 1
		dp[i] = min(dp[i-1], dp[(i+1)/2]) + 1
	}
	return dp[n]
}
