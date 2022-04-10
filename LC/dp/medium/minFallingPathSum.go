package medium

import "math"

// minFallingPathSum 931. 下降路径最小和
// fixme:自底向上
func minFallingPathSum(matrix [][]int) int {
	N := len(matrix)
	for r := N - 2; r >= 0; r-- {
		for c := 0; c < N; c++ {
			best := matrix[r+1][c]
			if c > 0 {
				best = min(best, matrix[r+1][c-1])
			}
			if c+1 < N {
				best = min(best, matrix[r+1][c+1])
			}
			matrix[r][c] += best
		}
	}

	ans := math.MaxInt64
	for i := range matrix[0] {
		ans = min(ans, matrix[0][i])
	}
	return ans
}
