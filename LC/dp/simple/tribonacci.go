package simple

// tribonacci 1137. 第 N 个泰波那契数
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	// 递归
	//dp := make([]int, n+1)
	//dp[0], dp[1], dp[2] = 0, 1, 1
	//for i := 3; i <= n; i++ {
	//	dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	//}
	//return dp[n]

	// 滚动数组
	p, q, r, s := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = s
		s = p + q + r
	}
	return s
}
