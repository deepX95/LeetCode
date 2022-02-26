package simple

// isRectangleOverlap 836. 矩形重叠
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 法1
	if rec1[0] == rec1[2] || rec1[1] == rec1[3] || rec2[0] == rec2[2] || rec2[1] == rec2[3] {
		return false
	}
	return !(rec1[2] <= rec2[0] ||   // left
		rec1[3] <= rec2[1] ||   // bottom
		rec1[0] >= rec2[2] ||   // right
		rec1[1] >= rec2[3])    // top


	// 法2
	//return min(rec1[2], rec2[2]) > max(rec1[0], rec2[0]) &&
	//		min(rec1[3], rec2[3]) > max(rec1[1], rec2[1])
}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
