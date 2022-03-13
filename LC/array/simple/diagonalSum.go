package simple

// diagonalSum 1572. 矩阵对角线元素的和
func diagonalSum(mat [][]int) int {
	// 逐行
	//sum := 0
	//n := len(mat)
	//for i := 0; i < n; i++ {
	//	sum += mat[i][i] + mat[i][n-1-i]
	//}
	//
	//return sum - mat[n/2][n/2]* (n & 1)

	sum := 0
	n := len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				sum += mat[i][j]
			}
		}
	}
	return sum

}
