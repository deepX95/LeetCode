package medium

// uniquePathsWithObstacles 63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range obstacleGrid {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			break
		}
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
			break
		}
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]

	// 滚动数组



}
