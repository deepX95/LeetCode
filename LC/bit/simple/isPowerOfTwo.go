package simple

// isPowerOfTwo 231. 2 的幂
// fixme:2的幂 则满足 n&(n-1)==0
func isPowerOfTwo(n int) bool {
	// 取模运算
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	return n == 1

	// 位运算
	//return n > 0 && n&(n-1) == 0 && n%3 == 1
	//return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0

}
