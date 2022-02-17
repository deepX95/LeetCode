package simple

// ClimbStairs 70. 爬楼梯
func ClimbStairs(n int) int {
	// 记忆搜索
	//dp:=make([]int,n+1)
	//dp[0],dp[1]=1,1
	//for i:=2;i<=n;i++{
	//	dp[i]=dp[i-1]+dp[i-2]
	//}
	//return dp[n]

	// 递归 超时
	if n<=3{
		return n
	}
	return ClimbStairs(n-1)+ClimbStairs(n-2)
}
