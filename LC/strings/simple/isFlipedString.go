package simple

// isFlipedString 面试题 01.09. 字符串轮转
func isFlipedString(s1 string, s2 string) bool {
	// 库函数
	//return len(s1) == len(s2) && strings.Contains(s2+s2, s1)

	// 位运算
	lens1, lens2, xor := len(s1), len(s2), 0
	if lens1 != lens2 {
		return false
	}
	for i, v := range s1 {
		xor ^= int(v) ^ int(s2[i])
	}
	return xor == 0

}
