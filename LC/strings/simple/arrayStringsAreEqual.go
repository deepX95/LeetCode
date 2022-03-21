package simple

// arrayStringsAreEqual 1662. 检查两个字符串数组是否相等
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	// 库函数
	//return strings.Join(word1, "") == strings.Join(word2, "")

	// 模拟
	var i, j, m, n int

	l1, l2 := len(word1), len(word2)

	for {
		if i >= l1 || m >= l2 {
			break
		}
		if word1[i][j] != word2[m][n] {
			return false
		}
		// 定位下一个
		j++
		if j >= len(word1[i]) {
			i++
			j = 0
		}
		n++
		if n >= len(word2[m]) {
			m++
			n = 0
		}
	}
	return i == l1 && m == l2 && j == 0 && n == 0

}
