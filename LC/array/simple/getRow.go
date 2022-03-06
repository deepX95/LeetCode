package simple

// getRow 119. 杨辉三角 II
// fixme:滚动数组
func getRow(rowIndex int) []int {
	// 递推
	//rows := make([][]int, rowIndex+1)
	//for i := range rows {
	//	rows[i] = make([]int, i+1)
	//	rows[i][0] = 1
	//	rows[i][i] = 1
	//	for j := 1; j < i; j++ {
	//		rows[i][j] = rows[i-1][j] + rows[i-1][j-1]
	//	}
	//}
	//return rows[rowIndex]

	// 滚动数组
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row

	// 线性递推
	//row := make([]int, rowIndex+1)
	//row[0] = 1
	//for i := 1; i <= rowIndex; i++ {
	//	row[i] = row[i-1] * (rowIndex - i + 1) / i
	//}
	//return row
}
