package simple

// hasAlternatingBits  693. 交替位二进制数
func hasAlternatingBits(n int) bool {
	// 0000 0101
	// 0000 1010

	// 0000 1010
	// 0000 0101

	// 模拟
	//cur, last := 0, -1
	//for i := 1; i < 32 && n > 0; i++ {
	//	cur = n & 1
	//	if cur == last {
	//		return false
	//	}
	//	last = cur
	//	n >>= 1
	//}
	//return true

	// fixme:题解
	// https://leetcode-cn.com/problems/binary-number-with-alternating-bits/solution/binary-number-with-alternating-bits-by-ikaruga/
	n = n ^ (n >> 1)
	return (n & (n + 1)) == 0
}
