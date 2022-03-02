package simple

// isPowerOfTwo 231. 2 的幂
// fixme:2的幂 则满足 n&(n-1)==0
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
