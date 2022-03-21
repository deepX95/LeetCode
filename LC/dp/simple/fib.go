package simple

// fib 509. 斐波那契数
// fixme:https://leetcode-cn.com/problems/fibonacci-number/solution/dong-tai-gui-hua-tao-lu-xiang-jie-by-labuladong/
func fib(n int) int {
	// 记忆化搜索
	//if n<2{
	//	return n
	//}
	//dp:= make([]int,n+1)
	//dp[0],dp[1]=0,1
	//for i:=2;i<=n;i++{
	//	dp[i]=dp[i-1]+dp[i-2]
	//}
	//return dp[n]

	// 递归
	//if n < 2 {
	//	return n
	//}
	//return fib(n-1) + fib(n-2)

	// 滚动数组
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a

}
