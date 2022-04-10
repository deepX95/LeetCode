package medium

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	//obstacleGrid := [][]int{
	//	{0, 0, 0},
	//	{0, 1, 0},
	//	{0, 0, 0},
	//}

	obstacleGrid := [][]int{
		{1}, {0},
	}

	uniquePathsWithObstacles(obstacleGrid)
}

func Test_minPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	minPathSum(grid)
}

func Test_minFallingPathSum(t *testing.T){
	grid := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	minFallingPathSum(grid)
}